package admin

import "github.com/gin-gonic/gin"

//实现方法的继承
type IndexController struct {
}

//(con UserController):表示把当前的函数挂载到结构体中,实现子结构体可以访问到父类中
//UserIndex：提供给外部进行调用:admin.UserIndex
func (con IndexController) Index(c *gin.Context) {
	c.String(200, "后台首页")
}
