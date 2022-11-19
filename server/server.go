package server

import (
	"context"
	"fmt"
	"jdlv/engine"
	"jdlv/engine/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/grid", grid)
	r.GET("/start", start)
	r.GET("/stop", stop)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func grid(c *gin.Context) {
	c.JSON(200, models.CurrentGrid().String())
}

func start(c *gin.Context) {
	if err := engine.Start(); err != nil {
		c.JSON(500, fmt.Sprintf("Start error %s", err.Error()))
		return
	}

	c.JSON(200, "Start success")
}

func stop(c *gin.Context) {
	if err := engine.Stop(); err != nil {
		c.JSON(500, fmt.Sprintf("Stop error %s", err.Error()))
		return
	}

	c.JSON(200, "Stop success")
}
