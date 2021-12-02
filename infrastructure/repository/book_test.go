package repository

import (
	"strings"
	"testing"

	"github.com/vcycyv/bookshop/entity"
	"github.com/vcycyv/bookshop/infrastructure/mock"
)

func TestCreate(t *testing.T) {
	db := mock.CreateDB()
	InitDB(db)
	repo := NewBookRepo(db)
	name := "test"
	book := entity.Book{
		Name: name,
		User: "tester",
	}
	newBook, err := repo.Add(book)
	if err != nil {
		t.Errorf("failed to add a book: %v", err)
		return
	}

	if !strings.EqualFold(newBook.Name, name) {
		t.Errorf("The added book is not correct.")
		return
	}
}
