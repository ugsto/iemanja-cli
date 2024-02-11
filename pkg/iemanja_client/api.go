package iemanjaclient

import (
	"context"
	"net"
	"net/http"
	"strings"
)

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

type ListPostsResponse struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

type NewPostRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type CreatePostResponse struct {
	Post Post `json:"post"`
}

type GetPostResponse struct {
	Post Post `json:"post"`
}

type UpdatePostResponse struct {
	Post Post `json:"post"`
}

type ListTagsResponse struct {
	Tags  []Tag `json:"tags"`
	Total int   `json:"total"`
}

type NewTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse struct {
	Tag Tag `json:"tag"`
}

type GetTagResponse struct {
	Tag Tag `json:"tag"`
}

type UpdateTagResponse struct {
	Tag Tag `json:"tag"`
}

type APIError struct {
	Error string `json:"error"`
}
