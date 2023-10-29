package main

import (
	"net/http"

	"github.com/aokuyama/circle_scheduler-api/interfaces/http/route/event"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
	engine.GET("/v1/event/:path", event.FetchEvent)
	engine.Run(":3000")
}
