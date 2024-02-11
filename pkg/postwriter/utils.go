package postwriter

import "strings"

func RemoveLastCharacterIfLineBreak(content string) string {
	if strings.HasSuffix(content, "\n") {
		return content[:len(content)-1]
	}
	return content
}

func ParseTags(tagsStr string) []string {
	return strings.FieldsFunc(tagsStr, func(r rune) bool {
		return r == ',' || r == ' '
	})
}
