package router

import (
	"yummy/controller"
	"github.com/kataras/iris"
)

func AdminRoutes(c *controller.Controller, api iris.Party) {
	api.Get("/posts", c.AdminGetPosts)
	api.Get("/posts/create", c.AdminGetCreatePostPage)
	api.Get("/product", c.AdminGetPosts)
}