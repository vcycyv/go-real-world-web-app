package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vcycyv/blog/entity"
	"github.com/vcycyv/blog/representation"
)

func TestBook_Add(t *testing.T) {
	name := "sqlite"
	book := representation.Book{
		Name: name,
		User: "user_a",
	}
	newBook, _ := bookSvc.Add(book)
	existingBook, _ := bookSvc.Get(newBook.ID)
	assert.Equal(t, name, existingBook.Name)

	_ = db.Migrator().DropTable(&entity.Book{})
	_ = db.Migrator().CreateTable(&entity.Book{})
}
