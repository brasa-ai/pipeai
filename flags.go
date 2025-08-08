package main

import "github.com/urfave/cli/v2"

var commonFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "ask",
		Usage:    "generate a command from natural language without executing it",
		Category: "Commands:",
	},
	&cli.StringFlag{
		Name:     "act",
		Usage:    "generate and execute a command from natural language",
		Category: "Commands:",
	},
	&cli.StringFlag{Name: "llm", Usage: "override provider", Category: "Options:"},
	&cli.StringFlag{Name: "model", Usage: "override model", Category: "Options:"},
	&cli.StringFlag{Name: "key", Usage: "override API key", Category: "Options:"},
	&cli.BoolFlag{Name: "debug", Usage: "enable debug logging", Category: "Options:"},
}
