package util

import (
	"Blog/util/db"
	"Blog/util/log"
	"Blog/util/mc"
	"Blog/util/redis"
)

// GatherMetrics 收集一些被动指标
func GatherMetrics() {
	mc.GatherMetrics()
	redis.GatherMetrics()
	db.GatherMetrics()
}

// Reset all utils
func Reset() {
	log.Reset()
}

// Stop all utils
func Stop() {
}
