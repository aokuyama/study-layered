package main

import (
	"os"
	"time"

	"github.com/aokuyama/circle_scheduler-api/packages/infrastructure/http/gin/route"
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
			"PUT",
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

	route.Define(g)

	g.Run(":3000")
}
