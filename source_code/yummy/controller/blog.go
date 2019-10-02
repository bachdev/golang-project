package controller

import (
	"yummy/model"
	"strconv"

	"github.com/kataras/iris"
)

func (c *Controller) GetBlog(ctx iris.Context) {
	// Viết code ở đây
	ctx.View("/blog/index.html")
}

func (c *Controller) GetPostById(ctx iris.Context) {
	// Viết code ở đây
	ctx.View("/blog/post.html")
}

// Phần này để viết hàm GetShop trong router
func (c *Controller) GetShop(ctx iris.Context) {
	// Viết code ở đây
	var data []model.Render_Product
	//Hàm c.DB.Query sẽ trả về có bao nhiêu bản ghi và lỗi nên phải hứng cái lỗi. Cái gạch là ko cần hứng
	_, err := c.DB.Query(&data, "SELECT id, name, thumbnail, price, description, created_at, modify_at, status FROM shop.product")
	if err != nil {
		panic(err)
	}
	ctx.ViewData("Product", data)
	ctx.View("/blog/shop.html")
}

func (c *Controller) GetMenu(ctx iris.Context) {
	// Viết code ở đây

	ctx.View("/blog/menu.html")
}

func (c *Controller) GetDetailProduct(ctx iris.Context) {
	// Viết code ở đây
	id := ctx.Params().Get("id")
	new_Id, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	//Đang làm dở ở đây.
	var detail_Product model.Render_Product
	_, err = c.DB.Query(&detail_Product, "SELECT id, name, thumbnail, price, description, created_at, modify_at, status FROM shop.product WHERE product.id=?", new_Id)
	ctx.ViewData("detail_Product", detail_Product)
	ctx.View("/blog/detailProduct.html")
}
