package model

import "time"

type Product struct {
	// Tao schema blog
	TableName []byte `sql:"shop.product"`
	// ID bài viết
	Id int32 `sql:",pk"`
	// Tên sản phẩm
	Name string
	// Thumbnail
	Thumbnail string
	//Giá
	Price int32
	// Nội dung
	Description string
	// Thời gian tạo
	CreatedAt time.Time
	// Thời gian sửa đổi
	ModifyAt time.Time
	// Trạng thái:
	Status int32 `sql:",notnull"`
}

type Render_Product struct {
	Id int32
	// Tên sản phẩm
	Name string
	// Thumbnail
	Thumbnail string
	//Giá
	Price int32
	// Nội dung
	Description string
	// Thời gian tạo
	CreatedAt time.Time
	// Thời gian sửa đổi
	ModifyAt time.Time
	// Trạng thái:
	Status int32
}
