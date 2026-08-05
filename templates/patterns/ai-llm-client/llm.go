package llm

import (
	"context"
	"errors"
	"io"

	"github.com/sashabaranov/go-openai"
)

type Config struct {
	BaseURL string
	APIKey  string
	Model   string
}

type Client struct {
	client *openai.Client
	model  string
}

func NewClient(cfg Config) *Client {
	config := openai.DefaultConfig(cfg.APIKey)
	if cfg.BaseURL != "" {
		config.BaseURL = cfg.BaseURL
	}

	return &Client{
		client: openai.NewClientWithConfig(config),
		model:  cfg.Model,
	}
}

// GenerateStream starts a streaming chat completion request and yields responses to a channel
func (c *Client) GenerateStream(ctx context.Context, systemPrompt, userMessage string) (<-chan string, <-chan error) {
	respChan := make(chan string)
	errChan := make(chan error, 1)

	go func() {
		defer close(respChan)
		defer close(errChan)

		req := openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
				{Role: openai.ChatMessageRoleUser, Content: userMessage},
			},
			Stream: true,
		}

		stream, err := c.client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			errChan <- err
			return
		}
		defer stream.Close()

		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}
			if err != nil {
				errChan <- err
				return
			}

			if len(response.Choices) > 0 {
				respChan <- response.Choices[0].Delta.Content
			}
		}
	}()

	return respChan, errChan
}
