package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vcycyv/blog/domain"
	rep "github.com/vcycyv/blog/representation"
)

type postHandler struct {
	postService domain.PostInterface
	authService domain.AuthInterface
}

func NewPostHandler(postService domain.PostInterface, authService domain.AuthInterface) postHandler {
	return postHandler{
		postService,
		authService,
	}
}

func (s *postHandler) Add(c *gin.Context) {
	post := &rep.Post{}
	if err := c.ShouldBind(post); err != nil {
		appErr := &rep.AppError{
			Code:    http.StatusBadRequest,
			Message: Message.InvalidMessage,
		}
		_ = c.Error(appErr)
		return
	}

	token := s.authService.ExtractToken(c)
	post.User, _ = s.authService.GetUserFromToken(token)

	fullPath := c.Request.URL.Scheme
	log.Print(fullPath)

	rtnVal, err := s.postService.Add(*post)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, rtnVal)
}

func (s *postHandler) Get(c *gin.Context) {
	id := c.Param("id")
	post, err := s.postService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (s *postHandler) GetAll(c *gin.Context) {
	posts, err := s.postService.GetAll()
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (s *postHandler) Update(c *gin.Context) {
	id := c.Param("id")
	_, err := s.postService.Get(id)
	if err != nil {
		_ = c.Error(err)
		return
	}

	post := &rep.Post{}
	if err := c.ShouldBind(post); err != nil {
		_ = c.Error(&rep.AppError{
			Code:    http.StatusBadRequest,
			Message: Message.InvalidMessage,
		})
		return
	}

	post, err = s.postService.Update(*post)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (s *postHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := s.postService.Delete(id)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
