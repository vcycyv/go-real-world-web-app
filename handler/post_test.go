package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vcycyv/blog/entity"
	"github.com/vcycyv/blog/representation"
)

func TestPost_Add(t *testing.T) {
	var postName = "sqlite"
	post := addPost(postName)
	assert.Equal(t, postName, post.Name)

	_ = db.Migrator().DropTable(&entity.Post{})
	_ = db.Migrator().CreateTable(&entity.Post{})
}

func TestPost_Get(t *testing.T) {
	w := httptest.NewRecorder()

	var postName = "sqlite"
	post := addPost(postName)
	postName = post.Name

	uri := fmt.Sprintf("/posts/%s", post.ID)
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodGet,
		uri,
		nil,
	)
	r.GET("/posts/:id", postHdlr.Get)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	got := &representation.Post{}
	_ = json.Unmarshal(body, &got)
	assert.True(t, len(got.ID) > 0)
	assert.Equal(t, postName, got.Name)

	_ = db.Migrator().DropTable(&entity.Post{})
	_ = db.Migrator().CreateTable(&entity.Post{})
}

func TestPost_GetAll(t *testing.T) {
	var postName = "sqlite"
	post := addPost(postName)

	posts := getPosts()

	assert.Equal(t, len(posts), 1)
	assert.Equal(t, post.Name, posts[0].Name)

	_ = db.Migrator().DropTable(&entity.Post{})
	_ = db.Migrator().CreateTable(&entity.Post{})
}

func TestPost_Update(t *testing.T) {
	w := httptest.NewRecorder()

	var postName = "sqlite"
	var renamed = "sqlite_renamed"
	post := addPost(postName)

	post.Name = renamed
	reqBody, _ := json.Marshal(post)

	uri := fmt.Sprintf("/posts/%s", post.ID)
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodPut,
		uri,
		bytes.NewReader(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	r.PUT("/posts/:id", postHdlr.Update)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	got := &representation.Post{}
	_ = json.Unmarshal(body, &got)
	assert.Equal(t, renamed, got.Name)

	_ = db.Migrator().DropTable(&entity.Post{})
	_ = db.Migrator().CreateTable(&entity.Post{})
}

func TestPost_Delete(t *testing.T) {
	w := httptest.NewRecorder()

	var postName = "sqlite"
	post := addPost(postName)

	uri := fmt.Sprintf("/posts/%s", post.ID)
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodDelete,
		uri,
		nil,
	)

	r.DELETE("/posts/:id", postHdlr.Delete)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	got := &representation.Post{}
	_ = json.Unmarshal(body, &got)

	_ = db.Migrator().DropTable(&entity.Post{})
	_ = db.Migrator().CreateTable(&entity.Post{})
}

func addPost(name string) *representation.Post {
	w := httptest.NewRecorder()

	uri := "/posts"
	r := gin.Default()

	post := representation.Post{
		Name: name,
	}
	reqBody, _ := json.Marshal(post)
	req := httptest.NewRequest(
		http.MethodPost,
		uri,
		bytes.NewReader(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	r.POST(uri, postHdlr.Add)
	//r.GET(uri, postController.GetAll)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(body, &post)
	return &post
}

func getPosts() []*representation.Post {
	w := httptest.NewRecorder()

	uri := "/posts"
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodGet,
		uri,
		nil,
	)
	r.GET(uri, postHdlr.GetAll)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	var posts []*representation.Post
	_ = json.Unmarshal(body, &posts)
	return posts
}
