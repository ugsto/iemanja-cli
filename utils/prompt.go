package utils

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func emptyCompleter(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func PromptInput(message string) string {
	return prompt.Input(message, emptyCompleter)
}

func PromptWithDefault(message string, defaultValue string) string {
	promptMessage := fmt.Sprintf("%s [%s]: ", message, defaultValue)
	response := prompt.Input(promptMessage, emptyCompleter)

	if response == "" {
		return defaultValue
	}
	return response
}
