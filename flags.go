package main

import "github.com/urfave/cli/v2"

var commonFlags = []cli.Flag{
	&cli.StringFlag{Name: "llm", Usage: "override provider"},
	&cli.StringFlag{Name: "model", Usage: "override model"},
	&cli.StringFlag{Name: "key", Usage: "override API key"},
	&cli.StringFlag{Name: "ask", Usage: "natural-language request"},
	&cli.BoolFlag{Name: "evaluate", Usage: "print command instead of executing"},
}
