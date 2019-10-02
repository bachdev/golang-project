package router

import (
	"github.com/kataras/iris"
	"yummy/controller"
)

func RegisterRoute(c *controller.Controller, app *iris.Application) {
	party := app.Party("/admin").Layout("layout/admin/backend-layout.html")
	{
		AdminRoutes(c, party)
	}

	BlogRoutes(c, app)
}