package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//配置session的中间件

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//initMiddleware:配置路由中间件
	r.GET("/", func(c *gin.Context) {
		//设置sessions
		session := sessions.Default(c)
		//配置session的过期时间
		session.Options(sessions.Options{MaxAge: 3600 * 6}) //6小时=60*60*6
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
