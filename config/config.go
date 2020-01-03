package config

import (
	"fmt"
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
func GetPort() int {
	return viper.GetInt("main_port")
}

func GetReadTimeOut() time.Duration {
	return viper.GetDuration("read_timeout")
}

func GetWriteTimeOut() time.Duration {
	return viper.GetDuration("write_timeout")
}

func GetMySQLConfig() string {
	host := viper.GetString("host")
	mysqlUser := viper.GetString("mysql_user")
	mysqlPassword := viper.GetString("mysql_password")
	mysqlPort := viper.GetString("mysql_port")
	databaseName := viper.GetString("database_name")
	mysqlConfig := mysqlUser + ":" + mysqlPassword + "@tcp(" + host + mysqlPort + ")" + databaseName
	return mysqlConfig
}
