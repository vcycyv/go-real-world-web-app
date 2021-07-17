package repository

import (
	"strings"
	"testing"

	"github.com/vcycyv/blog/entity"
	"github.com/vcycyv/blog/infrastructure/mock"
)

func TestCreate(t *testing.T) {
	db := mock.CreateDB()
	InitDB(db)
	repo := NewPostRepo(db)
	name := "test"
	post := entity.Post{
		Name: name,
		User: "tester",
	}
	newPost, err := repo.Add(post)
	if err != nil {
		t.Errorf("failed to add a post: %v", err)
		return
	}

	if !strings.EqualFold(newPost.Name, name) {
		t.Errorf("The added post is not correct.")
		return
	}
}
