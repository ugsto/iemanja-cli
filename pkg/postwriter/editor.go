package postwriter

import (
	"log"
	"os"
	"os/exec"
)

func EditContent(filetype, content string) string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	tmpFile, err := os.CreateTemp("", "post-*."+filetype)
	if err != nil {
		log.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if content != "" {
		if err := os.WriteFile(tmpFile.Name(), []byte(content), 0644); err != nil {
			log.Fatalf("Error writing to temporary file: %v", err)
		}
	}

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

	return RemoveLastCharacterIfLineBreak(string(contentBytes))
}
