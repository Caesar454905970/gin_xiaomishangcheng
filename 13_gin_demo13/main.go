package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//配置session的中间件
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret11111"))
	// 设置 session 中间件，参数 mysession，指的是 session 的名字，也是 cookie 的名字
	// store 是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	//initMiddleware:配置路由中间件
	r.GET("/", func(c *gin.Context) {
		//设置sessions
		session := sessions.Default(c)
		session.Set("username", "成强")
		//保存sessions:给其他页面使用（必须调用）
		session.Save()

		c.String(200, "gin首页")
	})
	r.GET("/news", func(c *gin.Context) {
		//获取sessions
		session := sessions.Default(c)
		username := session.Get("username")

		c.String(200, "username=%v", username)
	})

}
