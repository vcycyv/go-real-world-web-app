package domain

import (
	"github.com/vcycyv/bookshop/entity"
	rep "github.com/vcycyv/bookshop/representation"
)

type BookRepository interface {
	Add(book entity.Book) (*entity.Book, error)
	Get(id string) (*entity.Book, error)
	GetAll() ([]*entity.Book, error)
	Update(book entity.Book) (*entity.Book, error)
	Delete(id string) error
}

type BookInterface interface {
	Add(book rep.Book) (*rep.Book, error)
	Get(id string) (*rep.Book, error)
	GetAll() ([]*rep.Book, error)
	Update(book rep.Book) (*rep.Book, error)
	Delete(id string) error
}
