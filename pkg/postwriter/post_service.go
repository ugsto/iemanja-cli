package postwriter

import (
	iemanjaclient "github.com/ugsto/iemanja-cli/pkg/iemanja_client"
)

func SavePost(client *iemanjaclient.APIClient, id *string, title, content string, tags []string) (*iemanjaclient.Post, error) {
	if id != nil {
		resp, err := client.UpdatePost(*id, iemanjaclient.NewPostRequest{Title: title, Content: content, Tags: tags})

		if err != nil {
			return nil, err
		}

		return &resp.Post, nil
	}
	rep, err := client.CreatePost(iemanjaclient.NewPostRequest{Title: title, Content: content, Tags: tags})

	if err != nil {
		return nil, err
	}

	return &rep.Post, nil
}
