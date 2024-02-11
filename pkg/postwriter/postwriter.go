package postwriter

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"log"

	iemanja "github.com/ugsto/iemanja-cli/pkg/cmd"
	"github.com/ugsto/iemanja-cli/third_party"
	"github.com/ugsto/iemanja-cli/utils"
)

func WritePost(client *third_party.APIClient, filetype string) {
	content := launchEditor(filetype)

	title := utils.PromptInput("Title: ")
	tagsStr := utils.PromptInput("Tags: ")
	tags := parseTags(tagsStr)

	fmt.Println("Writing post... | title: ", title, " | content: ", content, " | tags: ", tags)

	iemanja.CreatePost(client, title, content, tags)
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
	return string(contentBytes)
}

func parseTags(tagsStr string) []string {
	return strings.Split(tagsStr, ",")
}
