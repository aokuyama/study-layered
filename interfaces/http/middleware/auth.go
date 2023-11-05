package middleware

import (
	"net/http"
	"strings"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	var err error

	tokenString := parseBearerToken(ctx)
	if tokenString == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	token, err := user.NewAuthToken(*tokenString)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}

	jwt := middleware.NewJwt()

	id, err := jwt.AuthByToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}
	ctx.Set("AuthorizedUser", id)

	ctx.Next()
}

func parseBearerToken(ctx *gin.Context) *string {
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		ary := strings.Split(authorizationHeader, " ")
		if len(ary) == 2 {
			if ary[0] == "Bearer" {
				return &ary[1]
			}
		}
	}
	return nil
}
