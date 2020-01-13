package test

import (
	"log"
	"testing"
	"time"
)

/**
测试格式化时间
*/
func TestFormatTime(t *testing.T) {
	log.Println(util.FormatTime(time.Now()))
}

/**
测试解析时间字符串
*/
func TestParseTime(t *testing.T) {
	log.Println(util.ParseTime("2019-05-01 12:34:23"))
}

/**
测试获取当前的日期字符串
*/
func TestGetCurrentDateString(t *testing.T) {
	log.Println(util.GetCurrentDateString("/"))
	log.Println(util.GetCurrentDateString("-"))
}
