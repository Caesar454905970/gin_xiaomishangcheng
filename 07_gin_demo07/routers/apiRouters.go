package routers

import (
	"07_gin_demo07/controllers/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

//AdminRouters:首字母大写，共有的，可以被其他调用
func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		r.GET("/user", admin.UserController{}.UserIndex)
		r.GET("/user/add", admin.UserController{}.UserAdd)
		r.GET("/user/edit", admin.UserController{}.UserEdit)

	}
	fmt.Println(apiRouters)

}
