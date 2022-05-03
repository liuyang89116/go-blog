package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
)

func GetPostsByCategoryId(categoryId, page, pageSize int) (*models.CategoryResponse, error) {
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPageByCategoryId(categoryId, page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	total := dao.GetPostCountByCategoryId(categoryId)
	pageCount := (total-1)/10 + 1
	var pages []int // 1, 2, 3, 4, 5, ...
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}

	homeResponse := &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categories,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageCount,
	}
	categoryName := dao.GetCategoryNameById(categoryId)
	categoryResponse := &models.CategoryResponse{
		HomeResponse: homeResponse,
		CategoryName: categoryName,
	}

	return categoryResponse, nil
}
