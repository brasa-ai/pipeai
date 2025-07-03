package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/your-handle/pipeai/actions"
	"github.com/your-handle/pipeai/helpers"
	"github.com/your-handle/pipeai/services"
)

func main() {
	ctx := context.Background()
	app := &cli.App{
		Name:  "pipeai",
		Usage: "Generate & execute shell commands with Gemini, OpenAI or Ollama",
		Commands: []*cli.Command{
			{
				Name:   "setup",
				Usage:  "create or update ~/.pipeai/config.yaml",
				Action: func(*cli.Context) error { return actions.RunSetup() },
			},
		},
		Flags: commonFlags,
		Action: func(c *cli.Context) error {
			cfg, _ := services.Load()
			if v := c.String("llm"); v != "" {
				cfg.LLM = v
			}
			if v := c.String("model"); v != "" {
				cfg.Model = v
			}
			if v := c.String("key"); v != "" {
				cfg.Key = v
			}
			q := c.String("ask")
			if q == "" {
				return cli.Exit("--ask is required", 1)
			}
			return actions.RunAsk(ctx, cfg, q, c.Bool("evaluate"))
		},
	}
	if err := app.Run(os.Args); err != nil {
		helpers.Log.Fatal().Err(err).Msg("")
	}
}
