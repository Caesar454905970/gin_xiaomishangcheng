package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//统计一个请求的执行时间
func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")
	//调用该请求的剩余处理程序
	c.Next() //去执行回调函数外的语句
	fmt.Println("2-我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)

}
func initMiddlewareOne(ctx *gin.Context) {
	fmt.Println("initMiddlewareOne--1-执行中中间件")
	//表示终止调用该请求的剩余处理程序
	ctx.Abort()
	fmt.Println("initMiddlewareOne--2-执行中中间件")
}

func main() {
	r := gin.Default()
	//全局中间件
	r.Use(initMiddlewareOne, initMiddleware)
	//initMiddleware:配置路由中间件
	r.GET("/", initMiddleware, func(c *gin.Context) {
		c.String(200, "gin首页")
	})
	r.GET("/news", initMiddleware, func(c *gin.Context) {
		c.String(200, "新闻页面")
	})

}
