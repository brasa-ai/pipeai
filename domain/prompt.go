package domain

import "fmt"

func BuildPrompt(q string) string {
	return fmt.Sprintf(DefaultPrompt, q)
}
