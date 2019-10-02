package router

import (
	"sample-project/controller"
	"github.com/kataras/iris"
)

func AdminRoutes(c *controller.Controller, api iris.Party) {
	api.Get("/posts", c.AdminGetPosts)
	api.Get("/posts/create", c.AdminGetCreatePostPage)
}