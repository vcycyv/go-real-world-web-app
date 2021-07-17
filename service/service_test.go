package service

import (
	"github.com/vcycyv/blog/domain"
	"github.com/vcycyv/blog/infrastructure/mock"
	"github.com/vcycyv/blog/infrastructure/repository"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	postSvc domain.PostInterface
)

func init() {
	db = mock.CreateDB()
	repository.InitDB(db)
	postRepo := repository.NewPostRepo(db)
	postSvc = NewPostService(postRepo)
}
