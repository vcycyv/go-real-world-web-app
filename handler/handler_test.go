package handler

import (
	"github.com/vcycyv/bookshop/infrastructure/mock"
	"github.com/vcycyv/bookshop/infrastructure/repository"
	"github.com/vcycyv/bookshop/service"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	bookHdlr bookHandler
)

func init() {
	authService := mock.NewMockAuth()

	db = mock.CreateDB()
	repository.InitDB(db)
	bookRepo := repository.NewBookRepo(db)
	bookService := service.NewBookService(bookRepo)
	bookHdlr = NewBookHandler(bookService, authService)
}
