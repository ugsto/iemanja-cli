package third_party

import (
	"github.com/ugsto/iemanja-cli/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListPosts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "posts": [ { "id": "ifsppovu4c4ehwa553m3", "title": "Post title", "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "tags": [ { "id": "wbdqikortur234x3au3b", "name": "javascript" }, { "id": "qg6p00dfx2c03pbga4k4", "name": "typescript" } ], "created_at": "2024-02-11T01:26:04.018861335Z", "updated_at": "2024-02-11T01:26:04.018861335Z" } ], "total": 1 }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	response, err := api.ListPosts(1, 0)
	if err != nil {
		t.Fatalf("ListPosts returned an error: %v", err)
	}

	expected := ListPostsResponse{
		Posts: []model.Post{
			{
				ID:      "ifsppovu4c4ehwa553m3",
				Title:   "Post title",
				Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				Tags: []model.Tag{
					{ID: "wbdqikortur234x3au3b", Name: "javascript"},
					{ID: "qg6p00dfx2c03pbga4k4", Name: "typescript"},
				},
				CreatedAt: "2024-02-11T01:26:04.018861335Z",
				UpdatedAt: "2024-02-11T01:26:04.018861335Z",
			},
		},
		Total: 1,
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestCreatePost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "ifsppovu4c4ehwa553m3", "title": "Post title", "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "tags": [ { "id": "wbdqikortur234x3au3b", "name": "javascript" }, { "id": "qg6p00dfx2c03pbga4k4", "name": "typescript" } ], "created_at": "2024-02-11T01:26:04.018861335Z", "updated_at": "2024-02-11T01:26:04.018861335Z" }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	post := NewPostRequest{
		Title:   "Post title",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		Tags:    []string{"javascript", "typescript"},
	}

	response, err := api.CreatePost(post)
	if err != nil {
		t.Fatalf("CreatePost returned an error: %v", err)
	}

	expected := CreatePostResponse{
		Post: model.Post{
			ID:      "ifsppovu4c4ehwa553m3",
			Title:   "Post title",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Tags: []model.Tag{
				{ID: "wbdqikortur234x3au3b", Name: "javascript"},
				{ID: "qg6p00dfx2c03pbga4k4", Name: "typescript"},
			},
			CreatedAt: "2024-02-11T01:26:04.018861335Z",
			UpdatedAt: "2024-02-11T01:26:04.018861335Z",
		},
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestGetPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "ifsppovu4c4ehwa553m3", "title": "Post title", "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", "tags": [ { "id": "wbdqikortur234x3au3b", "name": "javascript" }, { "id": "qg6p00dfx2c03pbga4k4", "name": "typescript" } ], "created_at": "2024-02-11T01:26:04.018861335Z", "updated_at": "2024-02-11T01:26:04.018861335Z" }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	response, err := api.GetPost("ifsppovu4c4ehwa553m3")
	if err != nil {
		t.Fatalf("GetPost returned an error: %v", err)
	}

	expected := GetPostResponse{
		Post: model.Post{
			ID:      "ifsppovu4c4ehwa553m3",
			Title:   "Post title",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Tags: []model.Tag{
				{ID: "wbdqikortur234x3au3b", Name: "javascript"},
				{ID: "qg6p00dfx2c03pbga4k4", Name: "typescript"},
			},
			CreatedAt: "2024-02-11T01:26:04.018861335Z",
			UpdatedAt: "2024-02-11T01:26:04.018861335Z",
		},
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestUpdatePost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "ifsppovu4c4ehwa553m3", "title": "Post title - changed", "content": "Content Changed", "tags": [ { "id": "wbdqikortur234x3au3b", "name": "javascript" }, { "id": "qg6p00dfx2c03pbga4k4", "name": "typescript" } ], "created_at": "2024-02-11T01:26:04.018861335Z", "updated_at": "2024-02-11T01:26:04.018861335Z" }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	postID := "ifsppovu4c4ehwa553m3"
	post := NewPostRequest{
		Title:   "Post title - changed",
		Content: "Content Changed",
		Tags:    []string{"javascript", "typescript"},
	}

	response, err := api.UpdatePost(postID, post)
	if err != nil {
		t.Fatalf("UpdatePost returned an error: %v", err)
	}

	expected := UpdatePostResponse{
		Post: model.Post{
			ID:      "ifsppovu4c4ehwa553m3",
			Title:   "Post title - changed",
			Content: "Content Changed",
			Tags: []model.Tag{
				{ID: "wbdqikortur234x3au3b", Name: "javascript"},
				{ID: "qg6p00dfx2c03pbga4k4", Name: "typescript"},
			},
			CreatedAt: "2024-02-11T01:26:04.018861335Z",
			UpdatedAt: "2024-02-11T01:26:04.018861335Z",
		},
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestDeletePost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	err := api.DeletePost("ifsppovu4c4ehwa553m3")
	if err != nil {
		t.Fatalf("DeletePost returned an error: %v", err)
	}
}

func TestListTags(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "tags": [ { "id": "wbdqikortur234x3au3b", "name": "javascript" }, { "id": "qg6p00dfx2c03pbga4k4", "name": "typescript" } ], "total": 2 }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	response, err := api.ListTags(10, 0)
	if err != nil {
		t.Fatalf("ListTags returned an error: %v", err)
	}

	expected := ListTagsResponse{
		Tags: []model.Tag{
			{ID: "wbdqikortur234x3au3b", Name: "javascript"},
			{ID: "qg6p00dfx2c03pbga4k4", Name: "typescript"},
		},
		Total: 2,
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestCreateTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "wbdqikortur234x3au3b", "name": "javascript" }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	tag := NewTagRequest{Name: "javascript"}

	response, err := api.CreateTag(tag)
	if err != nil {
		t.Fatalf("CreateTag returned an error: %v", err)
	}

	expected := CreateTagResponse{
		Tag: model.Tag{
			ID:   "wbdqikortur234x3au3b",
			Name: "javascript",
		},
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestGetTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "wbdqikortur234x3au3b", "name": "javascript" }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	response, err := api.GetTag("javascript")
	if err != nil {
		t.Fatalf("GetTag returned an error: %v", err)
	}

	expected := GetTagResponse{
		Tag: model.Tag{
			ID:   "wbdqikortur234x3au3b",
			Name: "javascript",
		},
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestUpdateTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "id": "wbdqikortur234x3au3b", "name": "rust" }`))
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	tag := NewTagRequest{
		Name: "rust",
	}

	response, err := api.UpdateTag("javascript", tag)
	if err != nil {
		t.Fatalf("UpdateTag returned an error: %v", err)
	}

	expected := UpdateTagResponse{
		Tag: model.Tag{
			ID:   "wbdqikortur234x3au3b",
			Name: "rust",
		},
	}

	if !reflect.DeepEqual(*response, expected) {
		t.Errorf("Expected response %#v, got %#v", expected, *response)
	}
}

func TestDeleteTag(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	api := NewAPIClient(server.URL)

	err := api.DeleteTag("wbdqikortur234x3au3b")
	if err != nil {
		t.Fatalf("DeleteTag returned an error: %v", err)
	}
}
