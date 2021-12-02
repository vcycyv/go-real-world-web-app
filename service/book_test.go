package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vcycyv/bookshop/entity"
	"github.com/vcycyv/bookshop/representation"
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
