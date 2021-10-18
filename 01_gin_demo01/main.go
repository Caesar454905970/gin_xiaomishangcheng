package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "启动web服务")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(200, "这是一个Post请求，用户增加数据")
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这个一个Put请求 主要用于编辑数据")
	})
	r.DELETE("/del", func(c *gin.Context) {
		c.String(200, "这是一个DELETE用户删除数据")
	})

	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run(":8000") //启动一个web服务器
}
