package routers

import (
	"07_gin_demo07/controllers/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

//AdminRouters:首字母大写，共有的，可以被其他调用
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		r.GET("/", admin.IndexController{}.Index)

		r.GET("/article", admin.ArticleController{}.ArticleIndex)
		r.GET("/article/add", admin.ArticleController{}.ArticleAdd)
		r.GET("/article/edit", admin.ArticleController{}.ArticleEit)

	}
	fmt.Println(adminRouters)
}
