package user

import (
	"errors"
	"net/http"

	user_usecase "github.com/aokuyama/circle_scheduler-api/packages/application/create_user_by_password/usecase"
	auth_usecase "github.com/aokuyama/circle_scheduler-api/packages/application/user_create_auth_token/usecase"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/middleware"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var i user_usecase.CreateUserByPasswordInput
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
	f := user.UserFactoryImpl{}
	uu := user_usecase.New(f, r)
	userOut, err := uu.Invoke(&i)
	if err != nil {
		if errors.Is(err, errs.ErrBadParam) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
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
