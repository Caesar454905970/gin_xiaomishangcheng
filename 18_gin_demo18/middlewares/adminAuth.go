package middlewares

import "github.com/gin-gonic/gin"

//后端管理页面进行session读取，进行权限判断
func InitAdminAuthMiddleware(c *gin.Context) {
	//没有登录的用户，不能进入用户管理中心
}
