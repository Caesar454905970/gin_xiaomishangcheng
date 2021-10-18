package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	//`json:"content"`：返回给前端的字段（key）会小写：例如："title": "我是一个标题",
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//gin 相应数据
func main() {
	r := gin.Default()
	//配置模板文件的路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "首页")
	})

	r.GET("/json1", func(c *gin.Context) {
		//返回一个JSON数据,数据响应返回map数据类型
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好gin",
		})
	})
	//map[string]interface{}{} 封装成了 gin.h
	r.GET("/json2", func(c *gin.Context) {
		//返回一个JSON数据,数据响应返回map数据类型
		c.JSON(200, gin.H{
			"success": true,
			"msg":     "你好gin",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		//返回一个JSON数据,数据响应返回结构体
		c.JSON(200, a)
	})

	//响应Jsonp请求：可以传入回调函数:
	/*	例如：
		请求方式：http://localhost:8080/jsonp?callback=xxxx
		前端的返回结果：
		xxxx({
			"title": "我是一个标题",
				"desc": "描述",
				"content": "测试内容"
		})*/
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "我是一个标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		//返回一个JSONP数据,数据响应可以是回调函数
		c.JSONP(200, a)
	})

	//返回XML
	r.GET("/someXML", func(c *gin.Context) {
		//使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "成强"
		msg.Message = "gin"
		msg.Age = 12
		c.XML(http.StatusOK, msg)
	})
	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"news": "我是一个新闻的数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"goods": "我是一个商品的数据",
		})
	})
	//启动路由
	r.Run()
}
