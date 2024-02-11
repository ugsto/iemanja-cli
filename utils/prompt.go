package utils

import "github.com/c-bata/go-prompt"

func emptyCompleter(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func PromptInput(message string) string {
	return prompt.Input(message, emptyCompleter)
}
