package main

import (
	"go-blog/config"
	"go-blog/models"
	"log"
	"net/http"
	"text/template"
	"time"
)

type IndexData struct {
	Title string `json:"title"` // Title 必须大写，否则外面访问不到，网站没法显示
	Desc  string `json:"desc"`
}

func IsOdd(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	// 拿到当前路径
	path := config.Cfg.System.CurrentDir

	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	postList := path + "/template/layout/post-list.html"

	t.Funcs(template.FuncMap{"isOdd": IsOdd,
		"getNextName": GetNextName, "date": Date})

	// 因为首页有多个模板嵌套，解析的时候需要把它们都解析出来
	t, err := t.ParseFiles(path+"/template/index.html",
		home, header, footer, pagination, personal, postList)
	if err != nil {
		log.Println("Parse template error:", err)
	}

	// 必须定义页面上所有的数据
	categorys := []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	posts := []models.PostMore{
		{
			Pid:          1,
			Title:        "Bob's blog",
			Content:      "Demo",
			UserName:     "Bob",
			ViewCount:    123,
			CreateAt:     "2022-03-19",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	homeResponse := &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}

	// execute data
	t.Execute(w, homeResponse)
}

func main() {
	// app entry point
	// init web app, use http protocal
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// handle request
	http.HandleFunc("/", index)
	// 静态资源配置
	http.Handle("/resource/",
		http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
