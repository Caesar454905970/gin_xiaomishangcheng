package routers

import "github.com/gin-gonic/gin"

//AdminRouters:首字母大写，共有的，可以被其他调用
func AdminRoutersInit(r *gin.Engine) {
	adminRouters, _ := r.Group("/admin")
	{
		r.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		r.GET("/user", func(c *gin.Context) {
			c.String(200, "用户列表")
		})
		r.GET("/user/add", func(c *gin.Context) {
			c.String(200, "增加用户")
		})
		r.GET("/edit", func(c *gin.Context) {
			c.String(200, "修改用户")
		})

		r.GET("/article", func(c *gin.Context) {
			c.String(200, "新闻列表")
		})
	}
}
