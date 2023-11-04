package middleware

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClient struct{}

func NewJwt() *JwtClient {
	c := JwtClient{}
	return &c
}

func (c *JwtClient) CreateToken(i *user.UserID) (*user.AuthToken, error) {
	tokenString, err := GenerateToken(i.String())
	if err != nil {
		return nil, err
	}
	return user.NewAuthToken(*tokenString)
}

func (c *JwtClient) AuthByToken(i *user.UserID, t *user.AuthToken) error {
	return errors.New("unimplemented")
}

func GenerateToken(id string) (*string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if len(secretKey) <= 0 {
		panic("JWT_SECRET_KEY")
	}
	exp, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOUR"))
	if err != nil || exp <= 0 {
		panic("JWT_EXPIRE_HOUR")
	}

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
