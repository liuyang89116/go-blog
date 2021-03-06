package dao

import (
	"go-blog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println("Get all category error:", err)
		return nil, err
	}

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name,
			&category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("Read rows error:", err)
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryNameById(categoryId int) (categoryName string) {
	row := DB.QueryRow("select name from blog_category where cid=?", categoryId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&categoryName)

	return categoryName
}
