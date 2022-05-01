package models

import (
	"go-blog/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"` // 文章 ID
	Title      string    `json:"title"`
	Slug       string    `json:"slug"` // 自定义页面 path
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	CategoryId int       `json:"categoryId"`
	UserId     int       `json:"userId"`
	ViewCount  int       `json:"viewCount"`
	Type       int       `json:"type"` // 文章类型 0 普通，1 自定义文章
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}

type PostMore struct {
	Pid          int           `json:"pid"` // 文章 ID
	Title        string        `json:"title"`
	Slug         string        `json:"slug"` // 自定义页面 path
	Content      template.HTML `json:"content"`
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	UserId       int           `json:"userId"`
	UserName     string        `json:"userName"`
	ViewCount    int           `json:"viewCount"`
	Type         int           `json:"type"` // 文章类型 0 普通，1 自定义文章
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
