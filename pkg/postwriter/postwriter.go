package postwriter

import (
	"os"
	"os/exec"
	"strings"

	"log"

	"github.com/thoas/go-funk"
	iemanja "github.com/ugsto/iemanja-cli/pkg/cmd"
	iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"
	"github.com/ugsto/iemanja-cli/utils"
)

func WritePost(client *iemanjaclient.APIClient, filetype string) {
	content := launchEditor(filetype)

	title := utils.PromptInput("Title: ")
	tagsStr := utils.PromptInput("Tags: ")
	tags := parseTags(tagsStr)

	iemanja.CreatePost(client, title, content, tags)
}

func removeLastCharacterIfLineBreak(content string) string {
	if strings.HasSuffix(content, "\n") {
		return content[:len(content)-1]
	}
	return content
}

func launchEditor(filetype string) string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	tmpFile, err := os.CreateTemp("", "post-*."+filetype)
	if err != nil {
		log.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	cmd := exec.Command(editor, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error opening editor: %v", err)
	}

	contentBytes, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatalf("Error reading temporary file: %v", err)
	}

	content := removeLastCharacterIfLineBreak(string(contentBytes))

	return content
}

func parseTags(tagsStr string) []string {
	return funk.Filter(funk.Map(strings.Split(tagsStr, ","), strings.TrimSpace), func(tag string) bool { return tag != "" }).([]string)
}
