package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/your-handle/pipeai/helpers"
	"github.com/your-handle/pipeai/services"
)

func prompt(r *bufio.Reader, label, def string) string {
	fmt.Printf("%s [%s]: ", label, def)
	if v, err := r.ReadString('\n'); err == nil {
		if v = strings.TrimSpace(v); v != "" {
			return v
		}
	}
	return def
}

func RunSetup() error {
	cfg, err := services.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	r := bufio.NewReader(os.Stdin)
	cfg.LLM = prompt(r, "LLM provider (gemini/openai/ollama)", cfg.LLM)
	cfg.Key = prompt(r, "API key (skip for ollama)", cfg.Key)
	cfg.Model = prompt(r, "Model", cfg.Model)

	if err := services.Save(cfg); err != nil {
		return fmt.Errorf("failed to save config: %v", err)
	}

	helpers.Log.Info().Msg("Configuration saved ✔")
	return nil
}
