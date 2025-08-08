package domain

import (
	"fmt"
	"os"
)

type Config struct {
	LLM   string `yaml:"llm"`   // gemini | openai | ollama
	Model string `yaml:"model"` // gemini-pro, gpt-4o, llama3
	Key   string `yaml:"key"`   // API key (not needed for ollama)
}

const defaultPrompt = `You are an expert shell engineer.
Operating System: %s
Return ONLY the command (no markdown) that satisfies: "%s"`

func getOS() string {
	if os.PathSeparator == '\\' {
		return "Windows"
	}
	if _, err := os.Stat("/System/Library/CoreServices"); err == nil {
		return "macOS"
	}
	return "Linux"
}

func BuildPrompt(q string) string {
	return fmt.Sprintf(defaultPrompt, getOS(), q)
}
