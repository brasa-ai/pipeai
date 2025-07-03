package domain

type Config struct {
	LLM   string `yaml:"llm"`   // gemini | openai | ollama
	Model string `yaml:"model"` // gemini-pro, gpt-4o, llama3, …
	Key   string `yaml:"key"`   // API key (not needed for ollama)
}

const DefaultPrompt = `You are an expert Unix shell engineer.
Return ONLY the command (no markdown) that satisfies: "%s"`
