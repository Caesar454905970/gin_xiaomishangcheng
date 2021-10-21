package models

import "time"

//获取当前的日期
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

//获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}
