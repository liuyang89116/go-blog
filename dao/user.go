package dao

import "log"

func GetUserNameById(userId int) (userName string) {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
	}

	_ = row.Scan(&userName)
	return userName
}
