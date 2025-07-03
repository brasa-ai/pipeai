package actions

import (
	"context"
	"fmt"
	"strings"

	"github.com/your-handle/pipeai/domain"
	"github.com/your-handle/pipeai/helpers"
	"github.com/your-handle/pipeai/services"
)

// RunAsk generates a shell command with the LLM and executes (or prints) it.
func RunAsk(ctx context.Context, cfg *domain.Config, question string, dry bool) error {
	helpers.Log.Debug().Str("question", question).Msg("generating command") // <-- now helpers is used

	llm, err := services.NewLLM(ctx, cfg)
	if err != nil {
		return err
	}

	prompt := domain.BuildPrompt(question)
	resp, err := llm.Call(ctx, prompt)
	if err != nil {
		return err
	}

	cmd := strings.TrimSpace(resp)
	if dry {
		fmt.Println(cmd)
		return nil
	}
	return services.Run(cmd)
}
