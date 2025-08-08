package actions

import (
	"context"
	"fmt"
	"strings"

	"github.com/AxeByte/pipeai.axebyte/domain"
	"github.com/AxeByte/pipeai.axebyte/helpers"
	"github.com/AxeByte/pipeai.axebyte/services"
)

// RunAsk generates a shell command with the LLM and executes (or prints) it.
func RunAsk(ctx context.Context, cfg *domain.Config, question string, dry bool) error {
	helpers.Log.Debug().Msgf("Question = %q", question)

	llm, err := services.NewLLM(ctx, cfg)
	if err != nil {
		helpers.Log.Error().Msgf("Failed to initialize AI provider: %v", err)
		return fmt.Errorf("AI provider error: %v\nPlease check your configuration and ensure the service is available", err)
	}

	prompt := domain.BuildPrompt(question)
	resp, err := llm.Call(ctx, prompt)
	if err != nil {
		helpers.Log.Error().Msgf("Failed to generate command: %v", err)
		return fmt.Errorf("Command generation failed: %v\nPlease try again or check your AI provider's status", err)
	}

	// Clean up the command - remove markdown code blocks and extra whitespace
	cmd := strings.TrimSpace(resp)
	cmd = strings.TrimPrefix(cmd, "```bash\n")
	cmd = strings.TrimPrefix(cmd, "```sh\n")
	cmd = strings.TrimPrefix(cmd, "```\n")
	cmd = strings.TrimSuffix(cmd, "\n```")
	cmd = strings.TrimSpace(cmd)
	if cmd == "" {
		helpers.Log.Warn().Msg("AI returned empty command")
		return fmt.Errorf("No command was generated. Please try rephrasing your request")
	}

	if dry {
		helpers.Log.Info().Msgf(cmd)
		return nil
	}

	helpers.Log.Debug().Msg("Running command")
	if err := services.Run(cmd); err != nil {
		helpers.Log.Error().Msgf("Command execution failed: %v", err)
		return fmt.Errorf("Command execution failed: %v", err)
	}

	return nil
}
