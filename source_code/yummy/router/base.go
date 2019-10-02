package router

import (
	"github.com/kataras/iris"
	"sample-project/controller"
)

func RegisterRoute(c *controller.Controller, app *iris.Application) {
	party := app.Party("/admin").Layout("layout/admin-layout.html")
	{
		AdminRoutes(c, party)
	}

	BlogRoutes(c, app)
}