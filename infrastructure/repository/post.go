package repository

import (
	"errors"
	"fmt"
	"net/http"

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
	if err := m.db.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *postRepo) Get(id string) (*entity.Post, error) {
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

	return &data, err
}

func (m *postRepo) GetAll() ([]*entity.Post, error) {
	var posts []*entity.Post
	err := m.db.Find(&posts).Error
	if err != nil {
		return []*entity.Post{}, err
	}
	return posts, nil
}

func (m *postRepo) Update(post entity.Post) (*entity.Post, error) {
	err := m.db.Save(&post).Error

	return &post, err
}

func (m *postRepo) Delete(id string) error {
	tx := m.db.Begin()

	// maps := make(map[string]interface{})
	// maps["post_id"] = id
	// dataList, err := m.dataRepo.GetDataCollection(0, math.MaxInt32, maps)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// for i := range dataList {
	// 	err = m.dataRepo.DeleteData(dataList[i].ID)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }
	if err := tx.Where("id = ?", id).Delete(&entity.Post{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
