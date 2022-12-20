package server

import (
	"context"
	"fmt"
	"jdlv/engine"
	"jdlv/engine/models"
	"jdlv/server/controllers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PUT", "PATCH"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "Content-type"},
		MaxAge:        12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/grid", controllers.GetGrid)
	r.GET("/start", start)
	r.GET("/stop", stop)
	r.GET("/running", running)
	r.GET("/reinitialize", reinitialize)

	r.POST("/set_cells", controllers.SetCells)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func reinitialize(c *gin.Context) {
	models.CurrentGrid().Reinitialize()
	c.JSON(200, "success")
}

func running(c *gin.Context) {
	c.JSON(200, fmt.Sprintf("Running: %t", engine.Instance().IsRunning()))
}

func start(c *gin.Context) {
	if err := engine.Instance().Start(); err != nil {
		c.JSON(500, fmt.Sprintf("Start error %s", err.Error()))
		return
	}

	c.JSON(200, "Start success")
}

func stop(c *gin.Context) {
	if err := engine.Instance().Stop(); err != nil {
		c.JSON(500, fmt.Sprintf("Stop error %s", err.Error()))
		return
	}

	c.JSON(200, "Stop success")
}
