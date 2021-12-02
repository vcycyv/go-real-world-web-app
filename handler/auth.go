package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vcycyv/bookshop/domain"
)

type authHandler struct {
	authService domain.AuthInterface
}

func NewAuthHandler(authService domain.AuthInterface) authHandler {
	return authHandler{
		authService,
	}
}

func (s *authHandler) GetAuth(c *gin.Context) {
	type authInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	auth := authInfo{}
	if err := c.ShouldBind(&auth); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err := s.authService.Auth(auth.Username, auth.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	token, err := s.authService.GenerateToken(auth.Username, auth.Password)
	if err != nil {
		_ = c.AbortWithError(401, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "Logon suceeded.",
		"token": token,
	})
}

//might not useful
func (s *authHandler) CheckAuth(c *gin.Context) {
	token := s.authService.ExtractToken(c)
	var rtnVal bool
	if token == "" {
		rtnVal = false
	} else {
		claims, err := s.authService.ParseToken(token)
		if err != nil {
			rtnVal = false
		} else if time.Now().Unix() > claims.ExpiresAt {
			rtnVal = false
		} else {
			rtnVal = true
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"auth": rtnVal,
	})
}
