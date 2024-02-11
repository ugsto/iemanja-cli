package third_party

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/ugsto/iemanja-cli/model"
)

type ListPostsResponse struct {
	Posts []model.Post `json:"posts"`
	Total int          `json:"total"`
}

type NewPostRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type CreatePostResponse struct {
	Post model.Post `json:"post"`
}

type GetPostResponse struct {
	Post model.Post `json:"post"`
}

type UpdatePostResponse struct {
	Post model.Post `json:"post"`
}

type ListTagsResponse struct {
	Tags  []model.Tag `json:"tags"`
	Total int         `json:"total"`
}

type NewTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	Tag model.Tag `json:"tag"`
}

type GetTagResponse struct {
	Tag model.Tag `json:"tag"`
}

type UpdateTagResponse struct {
	Tag model.Tag `json:"tag"`
}

type APIClient struct {
	client  *http.Client
	baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	if strings.HasPrefix(baseURL, "unix") {
		socketPath := strings.TrimPrefix(baseURL, "unix://")
		return NewUnixApiClient(socketPath)
	}
	return NewTcpAPIClient(baseURL)
}

func NewTcpAPIClient(baseURL string) *APIClient {
	return &APIClient{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

func NewUnixApiClient(unixSocketPath string) *APIClient {
	return &APIClient{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", unixSocketPath)
				},
			},
		},
		baseURL: "http://unix",
	}
}

func (api *APIClient) ListPosts(limit, offset int) (*ListPostsResponse, error) {
	url := fmt.Sprintf("%s/api/v1/posts?limit=%d&offset=%d", api.baseURL, limit, offset)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Posts []model.Post `json:"posts"`
		Total int          `json:"total"`
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

	var newPost model.Post
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

	var post model.Post
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

	var updatedPost model.Post
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

	return nil
}

func (api *APIClient) ListTags(limit, offset int) (*ListTagsResponse, error) {
	url := fmt.Sprintf("%s/api/v1/tags?limit=%d&offset=%d", api.baseURL, limit, offset)
	resp, err := api.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Tags  []model.Tag `json:"tags"`
		Total int         `json:"total"`
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

	var newTag model.Tag
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

	var tag model.Tag
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

	var updatedTag model.Tag
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

	return nil
}
