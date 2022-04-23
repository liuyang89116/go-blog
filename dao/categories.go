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
