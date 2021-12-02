package domain

import (
	"github.com/vcycyv/bookshop/entity"
	rep "github.com/vcycyv/bookshop/representation"
)

type PostRepository interface {
	Add(post entity.Post) (*entity.Post, error)
	Get(id string) (*entity.Post, error)
	GetAll() ([]*entity.Post, error)
	Update(post entity.Post) (*entity.Post, error)
	Delete(id string) error
}

type PostInterface interface {
	Add(post rep.Post) (*rep.Post, error)
	Get(id string) (*rep.Post, error)
	GetAll() ([]*rep.Post, error)
	Update(post rep.Post) (*rep.Post, error)
	Delete(id string) error
}
