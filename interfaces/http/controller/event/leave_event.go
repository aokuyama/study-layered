package event

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/auth"
	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/response"
	"github.com/aokuyama/circle_scheduler-api/packages/application/user_leave_from_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/gin-gonic/gin"
)

func LeaveEvent(c *gin.Context) {
	id := auth.GetAuthorizedUser(c)

	if id.String() != c.Param("user_id") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	i := usecase.UserLeaveFromEventInput{
		EventID: c.Param("event_id"),
		UserID:  c.Param("user_id"),
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
	if response.HandleCommonError(c, err) {
		return
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
