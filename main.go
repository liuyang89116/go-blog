package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string `json:"title"` // Title 必须大写，否则外面访问不到，网站没法显示
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	data := IndexData{
		Title: "Learn Go",
		Desc:  "Keep learning, keep improving",
	}
	t := template.New("index.html")
	// 拿到当前路径
	path, _ := os.Getwd()

	home := path + "template/home.html"
	header := path + "template/layout/header.html"
	footer := path + "template/layout/footer.html"
	pagination := path + "template/layout/pagination.html"
	personal := path + "template/layout/personal.html"
	postList := path + "template/layout/post-list.html"
	// 因为首页有多个模板嵌套，解析的时候需要把它们都解析出来
	t, _ = t.ParseFiles(path+"/template/index.html",
		home, header, footer, pagination, personal, postList)
	// 必须定义页面上所有的数据

	// execute data
	t.Execute(w, data)
}

func main() {
	// app entry point
	// init web app, use http protocal
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// handle request
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
