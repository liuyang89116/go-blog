package dao

import (
	"go-blog/models"
	"log"
)

func GetAllPosts(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println("get all post error:", err)
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("scan post error:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetTotalPostCount() (count int) {
	row := DB.QueryRow("select count(1) from blog_post")
	err := row.Scan(&count)
	if err != nil {
		log.Println("Get total post count error:", err)
		return
	}
	return
}

func GetPostCountByCategoryId(cId int) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where category_id=?", cId)
	_ = row.Scan(&count)
	return count
}

func GetPostPageByCategoryId(categoryId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id=? limit ?,?",
		categoryId, page, pageSize)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("get post by category id error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
