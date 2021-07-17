package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type AuthInterface interface {
	Auth(user string, password string) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (*Claims, error)
	GetUserFromToken(token string) (string, error)
	ExtractToken(c *gin.Context) string
}
