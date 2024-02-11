package utils

import "github.com/ugsto/iemanja-cli/model"

func TagsToString(tags []model.Tag) []string {
	var tag_strings []string
	for _, tag := range tags {
		tag_strings = append(tag_strings, tag.Name)
	}
	return tag_strings
}
