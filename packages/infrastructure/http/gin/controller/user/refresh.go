package user

import (
	"net/http"

	auth_usecase "github.com/aokuyama/circle_scheduler-api/packages/application/user_create_auth_token/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/http/gin/middleware/auth"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	id := auth.GetAuthorizedUser(c)

	ar := middleware.NewJwt()
	au := auth_usecase.New(ar)
	ai := auth_usecase.UserCreateAuthTokenInput{
		UserId: id.String(),
	}
	authOut, err := au.Invoke(&ai)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    id.String(),
		"token": authOut.Token.String(),
	})
}
