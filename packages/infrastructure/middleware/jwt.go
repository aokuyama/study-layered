package middleware

import (
	"errors"
	"fmt"
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
	tokenString, err := generateToken(i.String())
	if err != nil {
		return nil, err
	}
	return user.NewAuthToken(*tokenString)
}

func (c *JwtClient) AuthByToken(t *user.AuthToken) (*user.UserID, error) {
	id, err := parseToken(t.String())
	if err != nil {
		return nil, err
	}
	return user.NewUserID(*id)
}

func generateToken(id string) (*string, error) {
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

func parseToken(tokenString string) (*string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if len(secretKey) <= 0 {
		panic("JWT_SECRET_KEY")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		i := claims["id"].(string)
		return &i, nil
	}
	return nil, errors.New("token invalid")
}
