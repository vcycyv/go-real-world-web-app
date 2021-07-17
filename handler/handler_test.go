package handler

import (
	"github.com/vcycyv/blog/infrastructure/mock"
	"github.com/vcycyv/blog/infrastructure/repository"
	"github.com/vcycyv/blog/service"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB
	postHdlr postHandler
)

func init() {
	authService := mock.NewMockAuth()

	db = mock.CreateDB()
	repository.InitDB(db)
	postRepo := repository.NewPostRepo(db)
	postService := service.NewPostService(postRepo)
	postHdlr = NewPostHandler(postService, authService)
}
