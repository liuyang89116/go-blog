package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	// 1. 页面 views; 2. 数据 api; 3. 静态资源
	http.HandleFunc("/", views.HTML.Index)
	// http://localhost:8080/c/1  1 参数, 分类的id
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.Handle("/resource/",
		http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
