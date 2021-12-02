package infrastructure

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"

	"github.com/vcycyv/bookshop/domain"
	rep "github.com/vcycyv/bookshop/representation"
)

type authService struct {
	conn         *ldap.Conn
	bindUsername string
	bindPassword string
	jwtSecret    string
}

func NewAuthService() domain.AuthInterface {
	rtnVal := &authService{}
	rtnVal.init()
	rtnVal.jwtSecret = "gin-vcycyv"

	return rtnVal
}

func (s *authService) init() {
	s.bindUsername = LDAPSetting.User
	s.bindPassword = LDAPSetting.Password

	l, err := ldap.DialURL(LDAPSetting.Url)
	if err != nil {
		log.Fatalf("Failed to dial: %s\n", err)
	}
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: s.bindUsername,
		Password: s.bindPassword,
	})
	s.conn = l
	if err != nil {
		log.Fatalf("Failed to bind: %s\n", err)
	}
}

func (s *authService) Auth(user string, password string) error {
	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		LDAPSetting.DC,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", user),
		[]string{"dn"},
		nil,
	)

	sr, err := s.conn.Search(searchRequest)
	if err != nil {
		return &rep.AppError{
			Code:    http.StatusUnauthorized,
			Message: fmt.Sprintf("ldap cannot search users: %v", err),
		}
	}

	if len(sr.Entries) != 1 {
		return &rep.AppError{
			Code:    http.StatusUnauthorized,
			Message: "User does not exist or too many entries returned",
		}
	}

	userdn := sr.Entries[0].DN

	// Bind as the user to verify their password
	err = s.conn.Bind(userdn, password)
	if err != nil {
		return &rep.AppError{
			Code:    http.StatusUnauthorized,
			Message: "username/password is incorrect",
		}
	}

	// Rebind as the read only user for any further queries
	err = s.conn.Bind(s.bindUsername, s.bindPassword)
	if err != nil {
		return &rep.AppError{
			Code:    http.StatusUnauthorized,
			Message: "rebind failed",
		}
	}

	return nil
}

// GenerateToken generate tokens used for auth
func (s *authService) GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := domain.Claims{
		Username: username, //EncodeMD5(username),
		Password: s.encodeMD5(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    s.jwtSecret,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(s.jwtSecret))

	return token, err
}

// ParseToken parsing token
func (s *authService) ParseToken(token string) (*domain.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*domain.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func (s *authService) GetUserFromToken(token string) (string, error) {
	claims, err := s.ParseToken(token)
	if err != nil {
		return "", err
	}

	return claims.Username, nil
}

func (s *authService) ExtractToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// EncodeMD5 md5 encryption
func (s *authService) encodeMD5(value string) string {
	m := md5.New()
	_, _ = m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
