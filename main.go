package main

import (
	"fmt"
	"github.com/ugsto/iemanja-cli/cmd"
	"github.com/ugsto/iemanja-cli/third_party"

	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&cmd.CLI)
	client := third_party.NewAPIClient(cmd.CLI.APIHost)

	switch ctx.Command() {
	case "list-posts":
		cmd.ListPosts(client, cmd.CLI.ListPosts.Limit, cmd.CLI.ListPosts.Offset)
	case "create-post":
		cmd.CreatePost(client, cmd.CLI.CreatePost.Title, cmd.CLI.CreatePost.Content, cmd.CLI.CreatePost.Tags)
	case "get-post":
		cmd.GetPost(client, cmd.CLI.GetPost.ID)
	case "update-post":
		cmd.UpdatePost(client, cmd.CLI.UpdatePost.ID, cmd.CLI.UpdatePost.Title, cmd.CLI.UpdatePost.Content, cmd.CLI.UpdatePost.Tags)
	case "delete-post":
		cmd.DeletePost(client, cmd.CLI.DeletePost.ID)
	case "list-tags":
		cmd.ListTags(client, cmd.CLI.ListTags.Limit, cmd.CLI.ListTags.Offset)
	case "create-tag":
		cmd.CreateTag(client, cmd.CLI.CreateTag.Name)
	case "get-tag":
		cmd.GetTag(client, cmd.CLI.GetTag.Name)
	case "update-tag":
		cmd.UpdateTag(client, cmd.CLI.UpdateTag.Name, cmd.CLI.UpdateTag.NewName)
	case "delete-tag":
		cmd.DeleteTag(client, cmd.CLI.DeleteTag.Name)
	default:
		fmt.Println("Command not recognized.")
	}
}
