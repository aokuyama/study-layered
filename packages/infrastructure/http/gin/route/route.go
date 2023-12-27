package route

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/http/gin/controller/event"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/http/gin/controller/user"
	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/http/gin/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Define(g *gin.Engine) {
	g.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	v1user := g.Group("/v1/user")
	v1user.Use(auth.Middleware)
	{
		v1user.GET("me", user.Me)
		v1user.GET("refresh", user.Refresh)
	}

	g.POST("/v1/user/signup", user.Signup)
	g.POST("/v1/user/auth", user.Auth)
	g.GET("/v1/e/:path", event.FetchEvent)

	v1event := g.Group("/v1/event")
	v1event.Use(auth.Middleware)
	{
		v1event.PUT(":event_id/member", event.JoinEvent)
		v1event.DELETE(":event_id/member/:user_id", event.LeaveEvent)
	}
}
