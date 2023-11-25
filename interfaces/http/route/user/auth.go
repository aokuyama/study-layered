package user

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/response"
	user_usecase "github.com/aokuyama/circle_scheduler-api/packages/application/auth_user_by_id_and_password/usecase"
	auth_usecase "github.com/aokuyama/circle_scheduler-api/packages/application/user_create_auth_token/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/middleware"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var i user_usecase.AuthUserByIDAndPasswordInput
	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := prisma.NewPrismaClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		p.Disconnect()
	}()

	r := prisma.NewUserRepositoryPrisma(p)
	uu := user_usecase.New(r)
	userOut, err := uu.Invoke(&i)
	if response.HandleCommonError(c, err) {
		return
	}

	ar := middleware.NewJwt()
	au := auth_usecase.New(ar)
	ai := auth_usecase.UserCreateAuthTokenInput{
		UserId: userOut.User.ID().String(),
	}
	authOut, err := au.Invoke(&ai)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    userOut.User.ID().String(),
		"token": authOut.Token.String(),
	})
}
