package user

import (
	"net/http"

	auth_usecase "github.com/aokuyama/circle_scheduler-api/packages/application/user_create_auth_token/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func Refresh(c *gin.Context) {
	i, ok := c.Get("AuthorizedUser")
	if !ok {
		panic("missing user")
	}

	ar := middleware.NewJwt()
	au := auth_usecase.New(ar)
	ai := auth_usecase.UserCreateAuthTokenInput{
		UserId: i.(*user.UserID).String(),
	}
	authOut, err := au.Invoke(&ai)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    i.(*user.UserID).String(),
		"token": authOut.Token.String(),
	})
}
