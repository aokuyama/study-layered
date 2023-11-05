package user

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	i, ok := c.Get("AuthorizedUser")
	if !ok {
		panic("missing user")
	}

	c.JSON(http.StatusOK, gin.H{
		"id": i.(*user.UserID).String(),
	})
}
