package routers

import (
	"catapi_project/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// This is for the main route
	// beego.Router("/", &controllers.MainController{})

	// This is for the /cat route
	web.Router("/", &controllers.CatController{}, "get:GetCatImage")
	web.Router("/vote", &controllers.CatController{}, "post:Vote")
}
