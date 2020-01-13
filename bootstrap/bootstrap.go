package bootstrap

import (
	"Blog/cache"
	"Blog/config"
	"Blog/cron"
	"Blog/dao"
	"Blog/ip_location"
	"Blog/kafka"
	"Blog/log"
	"Blog/mongo"
)

/**
系统初始化
*/
func Init() {
	config.Init()
	cache.Init()
	kafka.Init()
	mongo.Init()
	//dao层初始化
	dao.Init()

	log.Init()

	ip_location.Init()
	//初始化定时任务
	cron.Init()
}
