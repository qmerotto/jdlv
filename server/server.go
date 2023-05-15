package server

import (
	"context"
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

	gameGroup := r.Group("/game")
	gameGroup.POST("/start", controllers.StartGame)
	gameGroup.POST("/stop", controllers.StopGame)
	gameGroup.POST("/new", controllers.NewGame)
	gameGroup.Handle("GET", "/grid", func(c *gin.Context) {
		grid(c.Writer, c.Request)
	})
	jdlvGroup := r.Group("/jdlv")
	jdlvGroup.POST("/set_cells", controllers.SetCells)

	//r.GET("/grid", controllers.GetGrid)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
