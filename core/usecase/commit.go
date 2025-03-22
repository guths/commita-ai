package usecase

import (
	"context"
	"fmt"

	"github.com/guths/commita-ai/core/port"
)

type Summarize struct {
	ctx   context.Context
	aiApi port.AiClient
}

func NewSummarize(ctx context.Context, port port.AiClient) *Summarize {
	return &Summarize{
		aiApi: port,
		ctx:   ctx,
	}
}

func (c *Summarize) Create(stagedChanges []byte) (string, error) {
	res, err := c.aiApi.ChatCompletion(c.ctx, "Summarize this commit diff using min tokens possible, use bullet points", stagedChanges)

	if err != nil {
		return "", err
	}

	fmt.Println(res)

	return res, nil
}
