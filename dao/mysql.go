package dao

import (
	"Blog/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	getString := config.GetString("mysql.url")
	DB, err := gorm.Open("mysql", getString)
	if err != nil {
		log.Panicln("--Init MySQL Connection Error:", err, "--")
	}
	DB.LogMode(true)
	log.Println("--Init MySQL Connection Finished--")
}
