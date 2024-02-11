package iemanja

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/thoas/go-funk"
	iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"
	"github.com/ugsto/iemanja-cli/utils"
)

func ListPosts(client *iemanjaclient.APIClient, limit, offset int) {
	response, err := client.ListPosts(limit, offset)
	if err != nil {
		log.Fatalf("Error listing posts: %v", err)
	}

	w := csv.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Printf("Total Posts: %d\n\n", response.Total)

	if err := w.Write([]string{"ID", "Title", "Tags", "Created At", "Updated At"}); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, post := range response.Posts {
		tags := utils.JoinTags(funk.Get(post.Tags, "Name").([]string))
		record := []string{post.ID, post.Title, tags, post.CreatedAt, post.UpdatedAt}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
}

func CreatePost(client *iemanjaclient.APIClient, title, content string, tags []string) {
	post := iemanjaclient.NewPostRequest{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
	response, err := client.CreatePost(post)
	if err != nil {
		log.Fatalf("Error creating post: %v", err)
	}
	fmt.Printf("Post created successfully:\n\nID: %s,\nTitle: %s\n", response.Post.ID, response.Post.Title)
}

func GetPost(client *iemanjaclient.APIClient, id string) {
	response, err := client.GetPost(id)
	if err != nil {
		log.Fatalf("Error getting post: %v", err)
	}
	fmt.Printf("Post retrieved successfully:\n\nID: %s,\nTitle: %s\nContent: %s\nTags: %s\nCreated At: %s\nUpdated At: %s\n", response.Post.ID, response.Post.Title, response.Post.Content, utils.JoinTags(funk.Get(response.Post.Tags, "Name").([]string)), response.Post.CreatedAt, response.Post.UpdatedAt)
}

func UpdatePost(client *iemanjaclient.APIClient, id, title, content string, tags []string) {
	post := iemanjaclient.NewPostRequest{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
	response, err := client.UpdatePost(id, post)
	if err != nil {
		log.Fatalf("Error updating post: %v", err)
	}
	fmt.Printf("Post updated successfully:\n\nID: %s,\nTitle: %s\nTags: %s\nCreated At: %s\nUpdated At: %s\n", response.Post.ID, response.Post.Title, utils.JoinTags(funk.Get(response.Post.Tags, "Name").([]string)), response.Post.CreatedAt, response.Post.UpdatedAt)
}

func DeletePost(client *iemanjaclient.APIClient, id string) {
	err := client.DeletePost(id)
	if err != nil {
		log.Fatalf("Error deleting post: %v", err)
	}
	fmt.Printf("Post deleted successfully:\n\nID: %s\n", id)
}
