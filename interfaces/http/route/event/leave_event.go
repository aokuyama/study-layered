package event

import (
	"errors"
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/auth"
	"github.com/aokuyama/circle_scheduler-api/packages/application/user_leave_from_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/gin-gonic/gin"
)

func LeaveEvent(c *gin.Context) {
	id := auth.GetAuthorizedUser(c)

	i := usecase.UserLeaveFromEventInput{
		EventID: c.Param("id"),
		UserID:  id.String(),
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

	guest := []gin.H{}
	for _, g := range out.Event.Guest().Items() {
		guest = append(guest, gin.H{
			"id":     g.UserID().String(),
			"name":   g.Name(),
			"number": g.Number(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"event": gin.H{
			"id":    out.Event.ID().String(),
			"name":  out.Event.Name().String(),
			"guest": guest,
			// TODO いずれ置き換える
			"start_at": "2000-01-01 00:00:00",
			"remarks":  "",
		},
	})
}
