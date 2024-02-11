package utils

import "strings"

func JoinTags(tags []string) string {
	tagsStr := strings.Join(tags, "; ")
	if tagsStr == "" {
		tagsStr = "**No tags**"
	}

	return tagsStr
}
