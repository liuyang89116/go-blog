package main

import (
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

func init() {
	// 加载模板
	common.LoadTemplate()
}

func main() {
	// app entry point
	// init web app, use http protocol
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// router
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
