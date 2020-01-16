package dao

import (
	"Blog/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
)

var DB *gorm.DB

func InitMysql() {
	var err error
	mysqlURL := strings.Join([]string{config.GetString("mysql.user"), ":", config.GetString("mysql.password"), "@", config.GetString("mysql.protocol"), "(", config.GetString("mysql.host"), ":", config.GetString("mysql.port"), ")/", config.GetString("mysql.database"), "?", config.GetString("mysql.charset"), "&", config.GetString("mysql.timeSupport")}, "")
	DB, err = gorm.Open("mysql", mysqlURL)
	if err != nil {
		log.Panicln("--Init MySQL Connection Error:", err, "--")
	}
	DB.LogMode(true)
	log.Println("--Init MySQL Connection Finished--")
}
