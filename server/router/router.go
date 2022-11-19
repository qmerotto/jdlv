package router

import (
	"jdlv/server/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	setupRoutes()
}

func setupRoutes() {
	web.
		web.Router("/", &controllers.LoginController{}, "get:Log")
	web.Router("/grid", &controllers.GridController{}, "get:Get,post:Post")
}
