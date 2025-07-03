package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/your-handle/pipeai/helpers"
	"github.com/your-handle/pipeai/services"
)

// RunSetup interactively writes ~/.pipeai/config.yaml.
func RunSetup() error {
	cfg, _ := services.Load()
	r := bufio.NewReader(os.Stdin)

	ask := func(label, def string) string {
		fmt.Printf("%s [%s]: ", label, def)
		v, _ := r.ReadString('\n')
		v = strings.TrimSpace(v)
		if v == "" {
			return def
		}
		return v
	}

	cfg.LLM   = ask("LLM provider (gemini/openai/ollama)", cfg.LLM)
	cfg.Key   = ask("API key (skip for ollama)",           cfg.Key)
	cfg.Model = ask("Model",                               cfg.Model)

	if err := services.Save(cfg); err != nil {
		return err
	}
	helpers.Log.Info().Msg("configuration saved ✔")
	return nil
}
