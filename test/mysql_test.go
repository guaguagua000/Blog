package main

import (
	"Blog/config"
	"Blog/dao"
	"Blog/model"
	"database/sql"
)

func main() {
	config.Init()
	dao.InitMysql()
	DB := dao.DB
	user := model.User{
		Id:       0,
		Username: "meowovo",
		Password: "abcd123456",
		Nickname: "Meow",
		Avatar:   "",
		Sex: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
	}
	DB.NewRecord("Users")
	DB.Create(&user)
	DB.Close()
}
