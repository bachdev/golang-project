package controller

import "github.com/kataras/iris"

func (c *Controller) AdminGetPosts(ctx iris.Context) {
	// Viết code ở đây
	ctx.View("/admin/blog/index.html")
}

func (c *Controller) AdminGetCreatePostPage(ctx iris.Context) {
	// Viết code ở đây
	
	ctx.View("/admin/blog/create.html")
}
func (c *Controller) AdminGetProductPage(ctx iris.Context) {
	// Viết code ở đây

	ctx.View("/admin/product/index.html")
}
