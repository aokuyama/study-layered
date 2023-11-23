package auth

import (
	"net/http"
	"strings"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	var err error

	tokenString := parseBearerToken(c)
	if tokenString == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}

	token, err := user.NewAuthToken(*tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}

	jwt := middleware.NewJwt()

	id, err := jwt.AuthByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}
	setAuthorizedUser(c, id)

	c.Next()
}

func parseBearerToken(c *gin.Context) *string {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		arr := strings.Split(authorizationHeader, " ")
		if len(arr) == 2 {
			if arr[0] == "Bearer" {
				return &arr[1]
			}
		}
	}
	return nil
}

func setAuthorizedUser(c *gin.Context, id *user.UserID) {
	c.Set("AuthorizedUser", id)
}

func GetAuthorizedUser(c *gin.Context) *user.UserID {
	id, ok := c.Get("AuthorizedUser")
	if !ok {
		panic("missing user")
	}
	return id.(*user.UserID)
}
