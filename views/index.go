package views

import (
	"go-blog/common"
	"go-blog/config"
	"go-blog/models"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	// 页面展示假数据
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

	index.WriteData(w, homeResponse)
}
