package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"log"
)

func GetAllIndexInfo() (*models.HomeResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		log.Println("Service get all category error:", err)
		return nil, err
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
		Categorys: categories,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}

	return homeResponse, nil
}
