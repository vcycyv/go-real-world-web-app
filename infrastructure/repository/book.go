package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/vcycyv/bookshop/domain"
	"github.com/vcycyv/bookshop/entity"
	"github.com/vcycyv/bookshop/representation"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) domain.BookRepository {
	return &bookRepo{
		db,
	}
}

func (m *bookRepo) Add(book entity.Book) (*entity.Book, error) {
	logrus.Debugf("about to save a book %s", book.Name)
	if err := m.db.Create(&book).Error; err != nil {
		return nil, err
	}
	logrus.Debugf("book %s saved", book.Name)
	return &book, nil
}

func (m *bookRepo) Get(id string) (*entity.Book, error) {
	logrus.Debugf("about to get a book %s", id)
	var data entity.Book
	err := m.db.Where("id = ?", id).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &representation.AppError{
				Code:    http.StatusFound,
				Message: fmt.Sprintf("Book %s is not found.", id),
			}
		}
		return &entity.Book{}, err
	}

	logrus.Debugf("book %s retrieved", id)
	return &data, err
}

func (m *bookRepo) GetAll() ([]*entity.Book, error) {
	logrus.Debug("about to get all book")
	var books []*entity.Book
	err := m.db.Find(&books).Error
	if err != nil {
		return []*entity.Book{}, err
	}
	logrus.Debug("all book retrieved")
	return books, nil
}

func (m *bookRepo) Update(book entity.Book) (*entity.Book, error) {
	logrus.Debugf("about to update a book %s", book.Name)
	err := m.db.Select("name", "updated_at").Updates(&book).Error
	logrus.Debugf("book %s updated", book.Name)
	return &book, err
}

func (m *bookRepo) Delete(id string) error {
	logrus.Debugf("about to delete a book %s", id)
	tx := m.db.Begin()
	if err := tx.Where("id = ?", id).Delete(&entity.Book{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	logrus.Debugf("book %s deleted", id)
	return tx.Commit().Error
}
