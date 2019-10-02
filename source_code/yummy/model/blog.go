package model

import "time"

type Post struct {
	// Tao shema blog
	TableName []byte `sql:"blog.post"`
	// ID bài viết
	Id string `sql:",pk"`
	// Tiêu đề
	Title string 
	// Thumbnail
	Thumbnail string 
	// Nội dung
	Content string 
	// Tác giả
	CreatedBy string 
	// Thời gian tạo 
	CreatedAt time.Time
	// Ngày xuất bản
	PublishedAt time.Time 
	// Trạng thái: 0 - nháp, 1 - xuất bản 
	Status int32 `sql:",notnull"`
}
