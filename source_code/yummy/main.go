package main

import (
	"sample-project/config"
	"sample-project/controller"
	"sample-project/model"
	"sample-project/router"

	"github.com/kataras/iris"
)

func main() {
	// Khởi tạo controller
	c := controller.NewController()
	cfg := config.SetupConfig()
	c.Config = cfg

	// Kết nối database
	db := model.ConnectDB(cfg.Database.User, cfg.Database.Password, cfg.Database.Database, cfg.Database.Address)
	c.DB = db
	defer db.Close()

	// Cấu hình database
	model.SetupDatabase(db, cfg)

	// Khởi tạo app iris
	app := iris.Default()
	app.Logger().SetLevel("debug")

	// Serve file tĩnh
	app.StaticWeb("/resources", "./resources")

	// Đăng ký view template từ thư mục view
	tmpl := iris.HTML("./view", ".html")
	tmpl.Layout("layout/layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)

	// Khai báo route
	router.RegisterRoute(c, app)

	// Chạy
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
