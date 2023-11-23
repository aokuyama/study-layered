package event

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/auth"
	"github.com/aokuyama/circle_scheduler-api/packages/application/user_join_to_event/usecase"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/persistence/prisma"
	"github.com/gin-gonic/gin"
)

func JoinEvent(c *gin.Context) {
	id := auth.GetAuthorizedUser(c)

	b := struct {
		User struct {
			Name   string
			Number int
		}
	}{}
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	i := usecase.UserJoinToEventInput{
		EventID: c.Param("id"),
		UserID:  id.String(),
		Name:    b.User.Name,
		Number:  uint8(b.User.Number),
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
		panic(err)
	}

	guest := []gin.H{}
	for _, g := range out.Event.Guest().Items() {
		guest = append(guest, gin.H{
			"id":     g.UserID().String(),
			"name":   *g.Name(),
			"number": *g.Number(),
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
		"user": gin.H{
			"id":     out.Guest.UserID().String(),
			"name":   out.Guest.Name(),
			"number": out.Guest.Number(),
		},
	})
}
