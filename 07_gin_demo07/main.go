package main

import (
	"07_gin_demo07/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type userInfo struct {
	//首字母大写表示公有的
	Username string `json:"username" form:"username" `
	Password string `json:"password" form:"password" `
}
type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)

	return str1 + "-----" + str2
}

func main() {
	r := gin.Default()

	//配置自定义模板函数
	r.SetFuncMap(template.FuncMap{
		//注册模板函数
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//配置模板文件的路径,放在路由配置的前面:加载templates下的所有文件
	r.LoadHTMLGlob("templates/**/*")

	//静态资源，第一个参数是路由，第二个是本地的目录
	r.Static("/static", "./static")

	//注册路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	r.Run()

}
