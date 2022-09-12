package router

import (
	"jdlv/server/controllers"

	"github.com/beego/beego/v2/server/web"
)

func SetupRoutes() {
	web.Router(
		"/", &controllers.LoginController{},
	)
}
