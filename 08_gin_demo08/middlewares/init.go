package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(ctx *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())

	fmt.Println(ctx.Request.URL)
	//设置值
	ctx.Set("username", "张三")

	//定义一个goroutine统计日志
	cCp := ctx.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done！ in path " + cCp.Request.URL.Path)
	}()
}
