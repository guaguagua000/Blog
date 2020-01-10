package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
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

func GetMySQLInfo() string {
	host := viper.GetString("host")
	mysqlUser := viper.GetString("mysql_user")
	mysqlPassword := viper.GetString("mysql_password")
	mysqlPort := viper.GetString("mysql_port")
	databaseName := viper.GetString("database_name")
	mysqlConfig := mysqlUser + ":" + mysqlPassword + "@tcp(" + host + mysqlPort + ")" + databaseName
	return mysqlConfig
}
