package main

import (
	"net/http"
	"os"
	"time"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/controller/event"
	"github.com/aokuyama/circle_scheduler-api/interfaces/http/controller/user"
	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	o := os.Getenv("ALLOW_ORIGIN")
	if len(o) < 1 {
		panic("ALLOW_ORIGIN")
	}

	g := gin.Default()

	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			o,
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

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
		v1event.POST(":id/member", event.JoinEvent)
		v1event.DELETE(":id/member", event.LeaveEvent)
	}

	g.Run(":3000")
}
