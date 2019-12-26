package config

import (
	"fmt"
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
	return viper.GetString("port")
}
