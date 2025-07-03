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

// NewLLM returns a concrete LLM implementation based on cfg.LLM.
//
// Supported values: gemini | openai | ollama (case-insensitive).
func NewLLM(ctx context.Context, cfg *domain.Config) (llms.LLM, error) {
	switch strings.ToLower(cfg.LLM) {
	case "gemini", "googleai":
		// LangChainGo uses GEMINI_API_KEY env var, but passing WithAPIKey is fine.
		llm, err := googleai.New(
			ctx,
			googleai.WithAPIKey(cfg.Key),
			googleai.WithDefaultModel(cfg.Model),
		)
		return llm, err

	case "openai":
		if cfg.Key != "" {
			_ = os.Setenv("OPENAI_API_KEY", cfg.Key) // fallback for provider internals
		}
		return openai.New(openai.WithModel(cfg.Model))

	case "ollama":
		return ollama.New(
			ollama.WithModel(cfg.Model),
			ollama.WithServerURL("http://localhost:11434"),
		)

	default:
		return nil, fmt.Errorf("unsupported LLM provider %q", cfg.LLM)
	}
}
