package router

import (
	"yummy/controller"

	"github.com/kataras/iris"
)

func BlogRoutes(c *controller.Controller, api *iris.Application) {
	api.Get("/blog", c.GetBlog)
	api.Get("/blog/{id}", c.GetPostById)
	// Hàm viết sau đó viết trong controller
	api.Get("/shop", c.GetShop)
	api.Get("/product/{id}", c.GetDetailProduct)
	api.Get("/menu", c.GetMenu)
}
