package response

import (
	"errors"
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/gin-gonic/gin"
)

func HandleCommonError(c *gin.Context, err error) (ng bool) {
	if err == nil {
		return false
	}
	if errors.Is(err, errs.ErrBadParam) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "bad request",
		})
		return true
	}
	if errors.Is(err, errs.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
		return true
	}
	if errors.Is(err, errs.ErrConflict) {
		c.JSON(http.StatusConflict, gin.H{
			"msg": "conflict",
		})
		return true
	}
	if errors.Is(err, errs.ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	panic(err)
}
