package views

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	if err := r.ParseForm(); err != nil {
		log.Println("Cannot parse form from request:", err)
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10 // every page display 10 blogs

	homeResponse, err := service.GetAllIndexInfo(page, pageSize)
	if err != nil {
		log.Println("Get index response error:", err)
		panic(err)
	}

	index.WriteData(w, homeResponse)
}
