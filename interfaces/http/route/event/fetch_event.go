package event

import (
	"errors"
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/packages/application/show_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/gin-gonic/gin"
)

func FetchEvent(c *gin.Context) {
	i := usecase.ShowEventInput{
		Path: c.Param("path"),
	}

	p, err := prisma.NewPrismaClient()
	if err != nil {
		panic(err)
	}
	defer func() {
		p.Disconnect()
	}()

	cr := prisma.NewEventRepositoryPrisma(p)
	u := usecase.New(cr)
	out, err := u.Invoke(&i)
	if err != nil {
		if errors.Is(err, errs.ErrBadParam) {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "bad request",
			})
			return
		}
		if errors.Is(err, errs.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "not found",
			})
			return
		}
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   out.Event.ID().String(),
		"name": out.Event.Name().String(),
	})
}
