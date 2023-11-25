package user

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	id := auth.GetAuthorizedUser(c)

	c.JSON(http.StatusOK, gin.H{
		"id": id.String(),
	})
}
