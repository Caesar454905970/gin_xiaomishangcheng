package routers

import "github.com/gin-gonic/gin"

//AdminRouters:首字母大写，共有的，可以被其他调用
func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		r.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api接口")
		})
		r.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个api接口--userlist")
		})
		r.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个api接口--plist")
		})

	}
}
