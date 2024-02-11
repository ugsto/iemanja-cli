package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	iemanja "github.com/ugsto/iemanja-cli/pkg/cmd"
	iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"
	"github.com/ugsto/iemanja-cli/pkg/postwriter"
)

var CLI struct {
	APIHost string `help:"API host. Can be an HTTP URL or a Unix socket path." default:"unix:///tmp/iemanja.sock"`

	WritePost struct {
		ID       string `arg:"" optional:"" help:"ID of the post to retrieve. If empty, a new post will be created."`
		FileType string `name:"filetype" default:"md" help:"File type for the post content (default: md)."`
	} `cmd:"" help:"Write a new post using the default editor."`

	ListPosts struct {
		Limit  int `short:"l" default:"10" help:"Limit number of posts to retrieve."`
		Offset int `short:"o" default:"0" help:"Offset for posts retrieval."`
	} `cmd:"" help:"List posts."`

	CreatePost struct {
		Title   string   `short:"t" required:"" help:"Title of the post."`
		Content string   `short:"c" required:"" help:"Content of the post."`
		Tags    []string `short:"T" required:"" help:"Tags of the post."`
	} `cmd:"" help:"Create a new post."`

	GetPost struct {
		ID string `arg:"" help:"ID of the post to retrieve."`
	} `cmd:"" help:"Get a post."`

	UpdatePost struct {
		ID      string   `arg:"" help:"ID of the post to update."`
		Title   string   `short:"t" required:"" help:"Title of the post."`
		Content string   `short:"c" required:"" help:"Content of the post."`
		Tags    []string `short:"T" required:"" help:"Tags of the post."`
	} `cmd:"" help:"Update a post."`

	DeletePost struct {
		ID string `arg:"" help:"ID of the post to delete."`
	} `cmd:"" help:"Delete a post."`

	ListTags struct {
		Limit  int `short:"l" default:"10" help:"Limit number of tags to retrieve."`
		Offset int `short:"o" default:"0" help:"Offset for tags retrieval."`
	} `cmd:"" help:"List tags."`

	CreateTag struct {
		Name string `arg:"" required:"" help:"Name of the tag."`
	} `cmd:"" help:"Create a new tag."`

	GetTag struct {
		Name string `arg:"" required:"" help:"Name of the tag to retrieve."`
	} `cmd:"" help:"Get a tag."`

	UpdateTag struct {
		OriginalName string `arg:"" required:"" help:"Name of the tag to update."`
		Name         string `short:"n" required:"" help:"New name of the tag."`
	} `cmd:"" help:"Update a tag."`

	DeleteTag struct {
		Name string `arg:"" required:"" help:"Name of the tag to delete."`
	} `cmd:"" help:"Delete a tag."`
}

func main() {
	ctx := kong.Parse(&CLI)
	client := iemanjaclient.NewAPIClient(CLI.APIHost)

	switch ctx.Command() {
	case "write-post":
		postwriter.WritePost(client, CLI.WritePost.FileType, CLI.WritePost.ID)
	case "write-post <id>":
		postwriter.WritePost(client, CLI.WritePost.FileType, CLI.WritePost.ID)
	case "list-posts":
		iemanja.ListPosts(client, CLI.ListPosts.Limit, CLI.ListPosts.Offset)
	case "create-post":
		iemanja.CreatePost(client, CLI.CreatePost.Title, CLI.CreatePost.Content, CLI.CreatePost.Tags)
	case "get-post <id>":
		iemanja.GetPost(client, CLI.GetPost.ID)
	case "update-post <id>":
		iemanja.UpdatePost(client, CLI.UpdatePost.ID, CLI.UpdatePost.Title, CLI.UpdatePost.Content, CLI.UpdatePost.Tags)
	case "delete-post <id>":
		iemanja.DeletePost(client, CLI.DeletePost.ID)
	case "list-tags":
		iemanja.ListTags(client, CLI.ListTags.Limit, CLI.ListTags.Offset)
	case "create-tag <name>":
		iemanja.CreateTag(client, CLI.CreateTag.Name)
	case "get-tag <name>":
		iemanja.GetTag(client, CLI.GetTag.Name)
	case "update-tag <original-name>":
		iemanja.UpdateTag(client, CLI.UpdateTag.OriginalName, CLI.UpdateTag.Name)
	case "delete-tag <name>":
		iemanja.DeleteTag(client, CLI.DeleteTag.Name)
	default:
		fmt.Printf("Command not recognized: %s\n", ctx.Command())
	}
}
