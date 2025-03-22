package port

import "context"

type AiClient interface {
	ChatCompletion(ctx context.Context, prompt string, data []byte) (string, error)
}
