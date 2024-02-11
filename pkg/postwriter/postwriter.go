package postwriter

import (
	"log"
	"strings"

	"github.com/thoas/go-funk"
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

	post, err := SavePost(client, postID, postData.Title, postData.Content, postData.Tags)

	if err != nil {
		log.Fatalf("Error saving post: %v", err)
	}

	log.Printf("Post saved successfully:\n\nID: %s,\nTitle: %s\nTags: %s\nCreated At: %s\nUpdated At: %s\n", post.ID, post.Title, utils.JoinTags(funk.Get(post.Tags, "Name").([]string)), post.CreatedAt, post.UpdatedAt)
}
