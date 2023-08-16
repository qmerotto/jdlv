package server

import (
	"context"
	"jdlv/server/controllers"
	"jdlv/server/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(ctx context.Context) {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PUT", "PATCH", "POST"},
		AllowHeaders:  []string{"Origin", "Content-type"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "Content-type"},
		MaxAge:        12 * time.Hour,
	}))

	r.Use(middleware.Auth)

	gameGroup := r.Group("/game")
	gameGroup.Handle("GET", "/grid", func(c *gin.Context) {
		grid(c.Writer, c.Request)
	})

	gameGroup.POST("/new", controllers.NewGame)
	gameGroup.POST("/token", controllers.CreateGameToken)
	gameGroup.POST("/stop", controllers.StopGame)

	jdlvGroup := gameGroup.Group("/jdlv")
	jdlvGroup.POST("/cell", controllers.SetCell)

	//r.GET("/grid", controllers.GetGrid)*/

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
