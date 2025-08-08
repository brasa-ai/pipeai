package actions

import (
	"context"
	"testing"

	"github.com/AxeByte/pipeai.axebyte/domain"
)

func TestRunAsk(t *testing.T) {
	// Test with invalid configuration
	err := RunAsk(context.Background(), &domain.Config{
		LLM:   "invalid",
		Model: "invalid",
		Key:   "invalid",
	}, "test command", true)
	if err == nil {
		t.Error("Expected error with invalid configuration")
	}

	// Test with empty question
	err = RunAsk(context.Background(), &domain.Config{
		LLM:   "ollama",
		Model: "llama2",
	}, "", true)
	if err == nil {
		t.Error("Expected error with empty question")
	}
}
