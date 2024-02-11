package iemanja

import (
	"fmt"
	"log"
	"strings"

	"github.com/ugsto/iemanja-cli/third_party"
	"github.com/ugsto/iemanja-cli/utils"
)

func ListPosts(client *third_party.APIClient, limit, offset int) {
	response, err := client.ListPosts(limit, offset)
	if err != nil {
		log.Fatalf("Error listing posts: %v", err)
	}
	fmt.Printf("Total Posts: %d\n\n", response.Total)
	for _, post := range response.Posts {
		fmt.Printf("ID: %s, Title: %s, Tags: %s\n", post.ID, post.Title, strings.Join(utils.TagsToString(post.Tags), "; "))
	}
}

func CreatePost(client *third_party.APIClient, title, content string, tags []string) {
	post := third_party.NewPostRequest{
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

func GetPost(client *third_party.APIClient, id string) {
	response, err := client.GetPost(id)
	if err != nil {
		log.Fatalf("Error getting post: %v", err)
	}
	fmt.Printf("Post retrieved successfully:\n\nID: %s,\nTitle: %s\nContent: %s\nTags: %s\n", response.Post.ID, response.Post.Title, response.Post.Content, strings.Join(utils.TagsToString(response.Post.Tags), "; "))
}

func UpdatePost(client *third_party.APIClient, id, title, content string, tags []string) {
	post := third_party.NewPostRequest{
		Title:   title,
		Content: content,
		Tags:    tags,
	}
	response, err := client.UpdatePost(id, post)
	if err != nil {
		log.Fatalf("Error updating post: %v", err)
	}
	fmt.Printf("Post updated successfully:\n\nID: %s,\nTitle: %s\n", response.Post.ID, response.Post.Title)
}

func DeletePost(client *third_party.APIClient, id string) {
	err := client.DeletePost(id)
	if err != nil {
		log.Fatalf("Error deleting post: %v", err)
	}
	fmt.Printf("Post deleted successfully:\n\nID: %s\n", id)
}

func ListTags(client *third_party.APIClient, limit, offset int) {
	tags, err := client.ListTags(limit, offset)
	if err != nil {
		log.Fatalf("Error listing tags: %v", err)
	}
	fmt.Printf("Total Tags: %d\n\n", tags.Total)
	for _, tag := range tags.Tags {
		fmt.Printf("ID: %s, Name: %s\n", tag.ID, tag.Name)
	}
}

func CreateTag(client *third_party.APIClient, name string) {
	tag := third_party.NewTagRequest{Name: name}
	response, err := client.CreateTag(tag)
	if err != nil {
		log.Fatalf("Error creating tag: %v", err)
	}
	fmt.Printf("Tag created successfully:\n\nID: %s,\nName: %s\n", response.Tag.ID, response.Tag.Name)
}

func GetTag(client *third_party.APIClient, id string) {
	response, err := client.GetTag(id)
	if err != nil {
		log.Fatalf("Error getting tag: %v", err)
	}
	fmt.Printf("Tag retrieved successfully:\n\nID: %s,\nName: %s\n", response.Tag.ID, response.Tag.Name)
}

func UpdateTag(client *third_party.APIClient, id, name string) {
	tag := third_party.NewTagRequest{Name: name}
	response, err := client.UpdateTag(id, tag)
	if err != nil {
		log.Fatalf("Error updating tag: %v", err)
	}
	fmt.Printf("Tag updated successfully:\n\nID: %s,\nName: %s\n", response.Tag.ID, response.Tag.Name)
}

func DeleteTag(client *third_party.APIClient, id string) {
	err := client.DeleteTag(id)
	if err != nil {
		log.Fatalf("Error deleting tag: %v", err)
	}
	fmt.Printf("Tag deleted successfully:\n\nID: %s\n", id)
}
