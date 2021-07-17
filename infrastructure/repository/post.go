package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vcycyv/blog/domain"
	"github.com/vcycyv/blog/entity"
	"github.com/vcycyv/blog/representation"
	"gorm.io/gorm"
)

type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) domain.PostRepository {
	return &postRepo{
		db,
	}
}

func (m *postRepo) Add(post entity.Post) (*entity.Post, error) {
	logrus.Debugf("about to save a post %s", post.Name)
	if err := m.db.Create(&post).Error; err != nil {
		return nil, err
	}
	logrus.Debugf("post %s saved", post.Name)
	return &post, nil
}

func (m *postRepo) Get(id string) (*entity.Post, error) {
	logrus.Debugf("about to get a post %s", id)
	var data entity.Post
	err := m.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &representation.AppError{
				Code:    http.StatusFound,
				Message: fmt.Sprintf("Post %s is not found.", id),
			}
		}
		return &entity.Post{}, err
	}

	logrus.Debugf("post %s retrieved", id)
	return &data, err
}

func (m *postRepo) GetAll() ([]*entity.Post, error) {
	logrus.Debug("about to get all post")
	var posts []*entity.Post
	err := m.db.Find(&posts).Error
	if err != nil {
		return []*entity.Post{}, err
	}
	logrus.Debug("all post retrieved")
	return posts, nil
}

func (m *postRepo) Update(post entity.Post) (*entity.Post, error) {
	logrus.Debugf("about to update a post %s", post.Name)
	err := m.db.Save(&post).Error
	logrus.Debugf("post %s updated", post.Name)
	return &post, err
}

func (m *postRepo) Delete(id string) error {
	logrus.Debugf("about to delete a post %s", id)
	tx := m.db.Begin()
	if err := tx.Where("id = ?", id).Delete(&entity.Post{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	logrus.Debugf("post %s deleted", id)
	return tx.Commit().Error
}
