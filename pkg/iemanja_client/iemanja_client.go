package iemanjaclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func handleAPIError(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	var apiErr APIError
	if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	return fmt.Errorf("API error: %s", apiErr.Error)
}

func (api *APIClient) ListPosts(limit, offset int) (*ListPostsResponse, error) {
	url := fmt.Sprintf("%s/api/v1/posts?limit=%d&offset=%d", api.baseURL, limit, offset)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var result struct {
		Posts []Post `json:"posts"`
		Total int    `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &ListPostsResponse{Posts: result.Posts, Total: result.Total}, nil
}

func (api *APIClient) CreatePost(newpost NewPostRequest) (*CreatePostResponse, error) {
	url := fmt.Sprintf("%s/api/v1/posts", api.baseURL)
	postData, err := json.Marshal(newpost)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var newPost Post
	if err := json.NewDecoder(resp.Body).Decode(&newPost); err != nil {
		return nil, err
	}

	return &CreatePostResponse{Post: newPost}, nil
}

func (api *APIClient) GetPost(postID string) (*GetPostResponse, error) {
	url := fmt.Sprintf("%s/api/v1/posts/%s", api.baseURL, postID)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var post Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}

	return &GetPostResponse{Post: post}, nil
}

func (api *APIClient) UpdatePost(postID string, post NewPostRequest) (*UpdatePostResponse, error) {
	url := fmt.Sprintf("%s/api/v1/posts/%s", api.baseURL, postID)
	postData, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var updatedPost Post
	if err := json.NewDecoder(resp.Body).Decode(&updatedPost); err != nil {
		return nil, err
	}

	return &UpdatePostResponse{Post: updatedPost}, nil
}

func (api *APIClient) DeletePost(postID string) error {
	url := fmt.Sprintf("%s/api/v1/posts/%s", api.baseURL, postID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return err
	}

	return nil
}

func (api *APIClient) ListTags(limit, offset int) (*ListTagsResponse, error) {
	url := fmt.Sprintf("%s/api/v1/tags?limit=%d&offset=%d", api.baseURL, limit, offset)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var result struct {
		Tags  []Tag `json:"tags"`
		Total int   `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &ListTagsResponse{Tags: result.Tags, Total: result.Total}, nil
}

func (api *APIClient) CreateTag(newtag NewTagRequest) (*CreateTagResponse, error) {
	url := fmt.Sprintf("%s/api/v1/tags", api.baseURL)
	tagData, err := json.Marshal(newtag)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.Post(url, "application/json", bytes.NewBuffer(tagData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var newTag Tag
	if err := json.NewDecoder(resp.Body).Decode(&newTag); err != nil {
		return nil, err
	}

	return &CreateTagResponse{Tag: newTag}, nil
}

func (api *APIClient) GetTag(tagName string) (*GetTagResponse, error) {
	url := fmt.Sprintf("%s/api/v1/tags/%s", api.baseURL, tagName)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var tag Tag
	if err := json.NewDecoder(resp.Body).Decode(&tag); err != nil {
		return nil, err
	}

	return &GetTagResponse{Tag: tag}, nil
}

func (api *APIClient) UpdateTag(tagName string, tag NewTagRequest) (*UpdateTagResponse, error) {
	url := fmt.Sprintf("%s/api/v1/tags/%s", api.baseURL, tagName)
	tagData, err := json.Marshal(tag)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(tagData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return nil, err
	}

	var updatedTag Tag
	if err := json.NewDecoder(resp.Body).Decode(&updatedTag); err != nil {
		return nil, err
	}

	return &UpdateTagResponse{Tag: updatedTag}, nil
}

func (api *APIClient) DeleteTag(tagName string) error {
	url := fmt.Sprintf("%s/api/v1/tags/%s", api.baseURL, tagName)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := handleAPIError(resp); err != nil {
		return err
	}

	return nil
}
