package routers

import (
	"08_gin_demo08/controllers/itying"
	"fmt"
	"github.com/gin-gonic/gin"
)

//DefaultRoutersInit:首字母大写，共有的，可以被其他调用
func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		r.GET("/", itying.DefaultController{}.Index)
		r.GET("/news", itying.DefaultController{}.News)

	}
	fmt.Println(defaultRouters)

}
