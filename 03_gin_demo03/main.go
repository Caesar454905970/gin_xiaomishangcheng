package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	//配置模板文件的路径,放在路由配置的前面:加载templates下的所有文件
	r.LoadHTMLGlob("templates/**/*")

	//前台页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "首页",
			"score": 89,
			//定义一个切片
			"hobby":     []string{"吃饭", "睡觉", "写代码"},
			"testSlice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题11",
					Content: "新闻内容11",
				},
				&Article{
					Title:   "新闻标题22",
					Content: "新闻内容22",
				},
			},
		})
	})
	r.GET("/news", func(c *gin.Context) {
		//实例化结构体
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})

	//后台
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.GET("/admin/news", func(c *gin.Context) {
		//实例化结构体
		news := &Article{
			Title:   "后台新闻标题",
			Content: "后台新闻内容",
		}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "后台页面",
			"news":  news,
		})
	})
	r.Run()

}
