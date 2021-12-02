package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/vcycyv/bookshop/domain"
	rep "github.com/vcycyv/bookshop/representation"
)

type bookHandler struct {
	bookService domain.BookInterface
	authService domain.AuthInterface
}

func NewBookHandler(bookService domain.BookInterface, authService domain.AuthInterface) bookHandler {
	return bookHandler{
		bookService,
		authService,
	}
}

func (s *bookHandler) Add(c *gin.Context) {
	book := &rep.Book{}
	if err := c.ShouldBind(book); err != nil {
		appErr := &rep.AppError{
			Code:    http.StatusBadRequest,
			Message: Message.InvalidMessage,
		}
		_ = c.Error(appErr)
		return
	}
	logger.Debugf("Received request to add a book %s.", book.Name)

	token := s.authService.ExtractToken(c)
	book.User, _ = s.authService.GetUserFromToken(token)

	rtnVal, err := s.bookService.Add(*book)
	if err != nil {
		_ = c.Error(err)
		return
	}
	logger.Debugf("The book %s is added successfully.", rtnVal.Name)
	c.JSON(http.StatusCreated, rtnVal)
}

func (s *bookHandler) Get(c *gin.Context) {
	id := c.Param("id")
	book, err := s.bookService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (s *bookHandler) GetAll(c *gin.Context) {
	books, err := s.bookService.GetAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (s *bookHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_, err := s.bookService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	book := &rep.Book{}
	if err := c.ShouldBind(book); err != nil {
		_ = c.Error(&rep.AppError{
			Code:    http.StatusBadRequest,
			Message: Message.InvalidMessage,
		})
		return
	}
	logger.Debugf("Received request to add a book %s", book.Name)

	book, err = s.bookService.Update(*book)
	if err != nil {
		_ = c.Error(err)
		return
	}
	logger.Debugf("The book %s is updated successfully.", book.Name)
	c.JSON(http.StatusOK, book)
}

func (s *bookHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	logger.Debugf("Received request to delete a book %s.", id)
	err := s.bookService.Delete(id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	logger.Debugf("The book %s is deleted successfully.", id)
	c.JSON(http.StatusNoContent, nil)
}
