package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//DefaultRoutersInit:首字母大写，共有的，可以被其他调用
func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "default/index.html", gin.H{
				"msg": "你好",
			})
		})
		r.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})
	}

}
