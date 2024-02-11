package postwriter

import (
	"strings"

	"github.com/ugsto/iemanja-cli/pkg/iemanja_client"
	"github.com/ugsto/iemanja-cli/utils"
)

func WritePost(client *iemanjaclient.APIClient, filetype string, id *string) {
	postData := RetrieveOrInitializePost(client, id)
	postData.Content = EditContent(filetype, postData.Content)

	postData.Title = utils.PromptWithDefault("Title: ", postData.Title)
	tagsStr := utils.PromptWithDefault("Tags (comma-separated): ", strings.Join(postData.Tags, ", "))
	postData.Tags = ParseTags(tagsStr)

	SavePost(client, id, postData.Title, postData.Content, postData.Tags)
}
