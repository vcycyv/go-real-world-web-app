package service

import (
	"github.com/vcycyv/blog/assembler"
	"github.com/vcycyv/blog/domain"
	rep "github.com/vcycyv/blog/representation"
)

type bookService struct {
	bookRepo domain.BookRepository
}

func NewBookService(bookRepo domain.BookRepository) domain.BookInterface {
	return &bookService{
		bookRepo,
	}
}

func (s *bookService) Add(book rep.Book) (*rep.Book, error) {
	data, err := s.bookRepo.Add(*assembler.BookAss.ToData(book))
	if err != nil {
		return &rep.Book{}, err
	}
	return assembler.BookAss.ToRepresentation(*data), nil
}

func (s *bookService) Get(id string) (*rep.Book, error) {
	data, err := s.bookRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return assembler.BookAss.ToRepresentation(*data), nil
}

func (s *bookService) GetAll() ([]*rep.Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		return nil, err
	}

	rtnVal := []*rep.Book{}
	for _, book := range books {
		rtnVal = append(rtnVal, assembler.BookAss.ToRepresentation(*book))
	}
	return rtnVal, nil
}

func (s *bookService) Update(book rep.Book) (*rep.Book, error) {
	data, err := s.bookRepo.Update(*assembler.BookAss.ToData(book))
	if err != nil {
		return nil, err
	}

	return assembler.BookAss.ToRepresentation(*data), nil
}

func (s *bookService) Delete(id string) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
