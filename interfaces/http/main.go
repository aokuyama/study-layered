package main

import (
	"net/http"
	"os"
	"time"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/middleware"
	"github.com/aokuyama/circle_scheduler-api/interfaces/http/route/event"
	"github.com/aokuyama/circle_scheduler-api/interfaces/http/route/user"
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
			"POST",
			"GET",
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
	v1user.Use(middleware.AuthMiddleware)
	{
		v1user.GET("me", user.Me)
	}
	g.POST("/v1/user/signup", user.Signup)
	g.GET("/v1/event/:path", event.FetchEvent)
	g.Run(":3000")
}
