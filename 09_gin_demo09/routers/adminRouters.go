package routers

import (
	"08_gin_demo08/controllers/itying"
	"08_gin_demo08/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

//AdminRouters:首字母大写，共有的，可以被其他调用
func AdminRoutersInit(r *gin.Engine) {
	//使用中间件
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	//使用中间件
	adminRouters.Use()
	{
		r.GET("/", admin.IndexController{}.Index)
		r.GET("/article", admin.ArticleController{}.ArticleIndex)
		r.GET("/article/add", admin.ArticleController{}.ArticleAdd)
		r.GET("/article/edit", admin.ArticleController{}.ArticleEit)

	}
	fmt.Println(adminRouters)
}
