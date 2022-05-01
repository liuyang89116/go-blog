package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var DB *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("yang:x@tcp(localhost:3306)/test?parseTime=true")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("connect sql error!")
		panic(err)
	}

	// 最大空闲连接数，默认不配置，是 2 个最大空闲连接
	db.SetMaxIdleConns(5)
	// 最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	// 空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)

	err = db.Ping()
	if err != nil {
		log.Println("Ping failed for database")
		_ = db.Close()
		panic(err)
	}

	DB = db
}
