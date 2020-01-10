package mysql

import (
	"Blog/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func InitMysql() {
	Db, err := sql.Open("mysql", config.GetMySQLInfo())
	if err != nil {
		panic(err)
		return
	}
	defer Db.Close()
}
