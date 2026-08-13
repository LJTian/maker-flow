package llm

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sashabaranov/go-openai"
)

// LLMClient is the decoupled abstract LLM client interface.
type LLMClient interface {
	GenerateStream(ctx context.Context, systemPrompt, userMessage string) (<-chan string, <-chan error)
}

// Config defines settings for LLM client.
type Config struct {
	Mode    string // "openai" or "mock"
	BaseURL string
	APIKey  string
	Model   string
}

// ConfigFromEnv reads configuration from environment variables.
func ConfigFromEnv() Config {
	mode := os.Getenv("LLM_MODE")
	if mode == "" {
		if os.Getenv("OPENAI_API_KEY") != "" {
			mode = "openai"
		} else {
			mode = "mock"
		}
	}
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-4o-mini"
	}
	return Config{
		Mode:    mode,
		BaseURL: os.Getenv("OPENAI_BASE_URL"),
		APIKey:  os.Getenv("OPENAI_API_KEY"),
		Model:   model,
	}
}

// NewLLMClient creates an LLMClient based on config mode.
func NewLLMClient(cfg Config) (LLMClient, error) {
	switch cfg.Mode {
	case "openai":
		return NewOpenAIClient(cfg)
	case "mock":
		return NewMockLLMClient("Mock LLM response: Hello from local dev mode!"), nil
	default:
		return NewMockLLMClient("Mock LLM response: Hello from local dev mode!"), nil
	}
}

// --- OpenAIClient Implementation ---

type OpenAIClient struct {
	client *openai.Client
	model  string
}

func NewOpenAIClient(cfg Config) (*OpenAIClient, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("openai mode requires OPENAI_API_KEY")
	}
	config := openai.DefaultConfig(cfg.APIKey)
	if cfg.BaseURL != "" {
		config.BaseURL = cfg.BaseURL
	}
	return &OpenAIClient{
		client: openai.NewClientWithConfig(config),
		model:  cfg.Model,
	}, nil
}

func (c *OpenAIClient) GenerateStream(ctx context.Context, systemPrompt, userMessage string) (<-chan string, <-chan error) {
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

// --- MockLLMClient Implementation (Local Dev / Testing) ---

type MockLLMClient struct {
	response string
}

func NewMockLLMClient(response string) *MockLLMClient {
	return &MockLLMClient{response: response}
}

func (m *MockLLMClient) GenerateStream(ctx context.Context, systemPrompt, userMessage string) (<-chan string, <-chan error) {
	respChan := make(chan string)
	errChan := make(chan error, 1)

	go func() {
		defer close(respChan)
		defer close(errChan)

		// Simulate token streaming
		for _, char := range m.response {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
				respChan <- string(char)
				time.Sleep(5 * time.Millisecond)
			}
		}
	}()

	return respChan, errChan
}
