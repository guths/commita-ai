package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/guths/commita-ai/internal/config"
)

type openAiClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewOpenAiClient() *openAiClient {
	return &openAiClient{
		apiKey:     config.GetAPIKey(),
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

type requestPayload struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type responsePayload struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *openAiClient) ChatCompletion(ctx context.Context, prompt string, data []byte) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	payload := requestPayload{
		Model: "gpt-3.5-turbo",
		Messages: []message{
			{
				Role:    "system",
				Content: "You are an assistant that summarizes git diffs concisely. Try to detailed in macro what is being changed, try to not use to many words, you do not comment what is being changed in every line, just pass a whole idea of files and what is being doing",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("%s:\n\n%s", prompt, string(data)),
			},
		},
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get response from OpenAI API")
	}

	var res responsePayload
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Choices) == 0 {
		return "", errors.New("empty response from OpenAI")
	}

	return res.Choices[0].Message.Content, nil
}
