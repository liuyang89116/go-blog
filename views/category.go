package views

import (
	"errors"
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	// http://localhost:8080/c/1  1 参数, 分类的id
	path := r.URL.Path
	cIdStr := strings.TrimPrefix(path, "/c/")
	categoryId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("Cannot find category path"))
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Parse form error:", err)
		categoryTemplate.WriteError(w, errors.New("Category parse form error!"))
		return
	}

	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(categoryId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
