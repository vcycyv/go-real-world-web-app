package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vcycyv/blog/entity"
	"github.com/vcycyv/blog/representation"
)

func TestPost_Add(t *testing.T) {
	name := "sqlite"
	post := representation.Post{
		Name: name,
		User: "user_a",
	}
	newPost, _ := postSvc.Add(post)
	existingPost, _ := postSvc.Get(newPost.ID)
	assert.Equal(t, name, existingPost.Name)

	_ = db.Migrator().DropTable(&entity.Post{})
	_ = db.Migrator().CreateTable(&entity.Post{})
}
