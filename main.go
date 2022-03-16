package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"` // Title 必须大写，否则外面访问不到，网站没法显示
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := IndexData{
		Title: "Learn Go",
		Desc:  "Keep learning, keep improving",
	}
	jsonStr, _ := json.Marshal(data)
	w.Write(jsonStr)
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
