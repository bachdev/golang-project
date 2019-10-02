# Project mẫu framework Iris 
Project sử dụng Go Modules do vậy yêu cầu cài đặt Go từ version 1.11.
Lưu ý: Nếu đặt project trong GOPATH để sử dụng Go Modules chạy câu lệnh
```
export GO111MODULE=on
```
Project là một trang blog có các route sau:
- /blog: trang danh sách bài viết 
- /blog/{id}: trang chi tiết bài viết 
- /admin: các trang quản trị sẽ bắt đầu bằng /admin 
    + admin/posts/create: trang tạo bài viết 
    + admin/posts: trang danh sách bài viết 
    + admin/posts/{id}: trang chỉnh sửa bài viết 

Cấu trúc source code:
- config: thư mục chứa struct cấu hình 
- constant: thư mục khai báo hằng 
- controller: thư mục chứa các hàm xử lý. mỗi hàm ứng với một route 
- helper: thư mục chứa các file 
- model: thư mục chứa các file định nghĩa model tương tác với go-pg
- resources: thư mục chứa tài nguyên static (css, js, image,...)
- router: thư mục chứa các file khai báo route 
- view: thư mục chứa các file html template 
- main.go: file chương trình chính 
- config.local.json: file cấu hình hệ thống (kết nối database, ...)

