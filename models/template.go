package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Customer   TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		w.Write([]byte("error write data..."))
	}
}

func InitTemplate(templateDir string) HtmlTemplate {
	tp := readTemplate(
		[]string{"index", "category", "customer",
			"detail", "login", "pigeonhole", "writing"},
		templateDir)

	var htmlTemplate HtmlTemplate
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Customer = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]

	return htmlTemplate
}

func readTemplate(templates []string, templateDir string) []TemplateBlog {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)

		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		pagination := templateDir + "layout/pagination.html"
		personal := templateDir + "layout/personal.html"
		postList := templateDir + "layout/post-list.html"

		// map 模板中的 method
		t.Funcs(template.FuncMap{
			"isOdd":       IsOdd,
			"getNextName": GetNextName,
			"date":        Date})

		// 因为首页有多个模板嵌套，解析的时候需要把它们都解析出来
		t, err := t.ParseFiles(templateDir+viewName,
			home, header, footer, pagination, personal, postList)
		if err != nil {
			log.Println("Parse template error:", err)
		}

		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}

	return tbs
}

func IsOdd(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}
