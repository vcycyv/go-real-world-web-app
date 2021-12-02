package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/vcycyv/bookshop/domain"
)

type auth struct{}

func NewMockAuth() domain.AuthInterface {
	return &auth{}
}

func (s *auth) Auth(user string, password string) error {
	return nil
}

func (s *auth) GenerateToken(username, password string) (string, error) {
	return "", nil
}

func (s *auth) ParseToken(token string) (*domain.Claims, error) {
	return nil, nil
}

func (s *auth) GetUserFromToken(token string) (string, error) {
	return "user_a", nil
}

func (s *auth) ExtractToken(c *gin.Context) string {
	return ""
}
