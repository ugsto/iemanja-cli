package utils

import iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"

func TagsToString(tags []iemanjaclient.Tag) []string {
	var tag_strings []string
	for _, tag := range tags {
		tag_strings = append(tag_strings, tag.Name)
	}
	return tag_strings
}
