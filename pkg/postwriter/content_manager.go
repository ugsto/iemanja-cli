package postwriter

import (
	"log"

	"github.com/thoas/go-funk"
	iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"
)

type PostData struct {
	Content string
	Title   string
	Tags    []string
}

func RetrieveOrInitializePost(client *iemanjaclient.APIClient, id *string) PostData {
	if id != nil {
		post, err := client.GetPost(*id)
		if err != nil {
			log.Fatalf("Error retrieving post: %v", err)
		}
		return PostData{
			Content: post.Post.Content,
			Title:   post.Post.Title,
			Tags:    funk.Get(post.Post.Tags, "Name").([]string),
		}
	}
	return PostData{}
}
