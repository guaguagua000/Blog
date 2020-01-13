package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile("config.toml")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
	}
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}
