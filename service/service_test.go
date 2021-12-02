package service

import (
	"github.com/vcycyv/bookshop/domain"
	"github.com/vcycyv/bookshop/infrastructure/mock"
	"github.com/vcycyv/bookshop/infrastructure/repository"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	bookSvc domain.BookInterface
)

func init() {
	db = mock.CreateDB()
	repository.InitDB(db)
	bookRepo := repository.NewBookRepo(db)
	bookSvc = NewBookService(bookRepo)
}
