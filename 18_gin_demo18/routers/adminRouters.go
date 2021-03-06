package routers

import (
	"18_gin_demo18/controllers/admin"
	"18_gin_demo18/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitAdminAuthMiddleware中间件(先走鉴权中间件)
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	{
		adminRouters.GET("/", admin.MainController{}.Index)

		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
		adminRouters.GET("/loginOut", admin.LoginController{}.LoginOut)

		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

	}
}
