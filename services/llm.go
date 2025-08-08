package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"

	"github.com/your-handle/pipeai/domain"
)

func NewLLM(ctx context.Context, cfg *domain.Config) (llms.LLM, error) {
	if cfg.Model == "" {
		return nil, fmt.Errorf("model name is required")
	}

	switch strings.ToLower(cfg.LLM) {
	case "gemini", "googleai":
		if cfg.Key == "" {
			return nil, fmt.Errorf("API key is required for Gemini")
		}
		return googleai.New(ctx, googleai.WithAPIKey(cfg.Key), googleai.WithDefaultModel(cfg.Model))

	case "openai":
		if cfg.Key == "" {
			return nil, fmt.Errorf("API key is required for OpenAI")
		}
		os.Setenv("OPENAI_API_KEY", cfg.Key)
		return openai.New(openai.WithModel(cfg.Model))

	case "ollama":
		return ollama.New(ollama.WithModel(cfg.Model), ollama.WithServerURL("http://localhost:11434"))

	default:
		return nil, fmt.Errorf("unsupported LLM provider %q (use: gemini, openai, or ollama)", cfg.LLM)
	}
}
