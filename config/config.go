package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

//读取配置文件
func LoadBaseConfig() {

	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
	}
}

//端口号
func GetPort() string {
	return viper.GetString("main_port")
}

func GetReadTimeOut() time.Duration {
	return viper.GetDuration("read_timeout")
}

func GetWriteTimeOut() time.Duration {
	return viper.GetDuration("write_timeout")
}

func ConnectMySQL() {
	host := viper.GetString("host")
	mysqlUser := viper.GetString("mysql_user")
	mysqlPassword := viper.GetString("mysql_password")
	mysqlPort := viper.GetString("mysql_port")
	databaseName := viper.GetString("database_name")
	mysqlConfig := mysqlUser + ":" + mysqlPassword + "@tcp(" + host + mysqlPort + ")" + databaseName
	Db, err := sql.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
		return
	}
	defer Db.Close()
}
