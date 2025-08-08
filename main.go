package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"

	"github.com/AxeByte/pipeai.axebyte/actions"
	"github.com/AxeByte/pipeai.axebyte/helpers"
	"github.com/AxeByte/pipeai.axebyte/services"
)

func run(c *cli.Context) error {
	// Load and update config
	cfg, err := services.Load()
	if err != nil {
		return cli.Exit("Failed to load config: "+err.Error(), 1)
	}
	helpers.SetupLogger(cfg.LogLevel)

	if v := c.String("llm"); v != "" {
		cfg.LLM = v
	}
	if v := c.String("model"); v != "" {
		cfg.Model = v
	}
	if v := c.String("key"); v != "" {
		cfg.Key = v
	}

	// Set log level
	if c.Bool("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Handle commands
	q := c.String("ask")
	act := c.String("act")

	switch {
	case q == "" && act == "":
		return cli.Exit("\nError: Either --ask or --act is required\nExamples:\n  pipeai --ask \"list files\"  # Generate only\n  pipeai --act \"list files\"  # Generate and execute\n", 1)
	case q != "" && act != "":
		return cli.Exit("\nError: Cannot use both --ask and --act together\n", 1)
	case act != "":
		return actions.RunAsk(c.Context, cfg, act, false)
	default:
		return actions.RunAsk(c.Context, cfg, q, true)
	}
}

func main() {
	app := &cli.App{
		Name:   "pipeai",
		Usage:  "Generate & execute shell commands with Gemini, OpenAI or Ollama",
		Flags:  commonFlags,
		Action: run,
		Commands: []*cli.Command{{
			Name:   "setup",
			Usage:  "create or update ~/.pipeai/config.yaml",
			Action: func(*cli.Context) error { return actions.RunSetup() },
		}},
	}

	if err := app.Run(os.Args); err != nil {
		helpers.Log.Fatal().Err(err).Msg("")
	}
}
