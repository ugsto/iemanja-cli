package postwriter

import (
	"log"
	"strings"

	"github.com/ugsto/iemanja-cli/pkg/iemanja_client"
	"github.com/ugsto/iemanja-cli/utils"
)

func WritePost(client *iemanjaclient.APIClient, filetype string, id string) {
	postID := &id
	if id == "" {
		postID = nil
	}

	postData := RetrieveOrInitializePost(client, postID)
	postData.Content = EditContent(filetype, postData.Content)

	if postData.Content == "" {
		log.Fatal("Content cannot be empty")
	}

	postData.Title = utils.PromptWithDefault("Title: ", postData.Title)
	tagsStr := utils.PromptWithDefault("Tags (comma-separated): ", strings.Join(postData.Tags, ", "))
	postData.Tags = ParseTags(tagsStr)

	SavePost(client, postID, postData.Title, postData.Content, postData.Tags)
}
