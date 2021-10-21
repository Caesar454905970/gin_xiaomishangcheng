package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//实现方法的继承
type IndexController struct {
}

//(con UserController):表示把当前的函数挂载到结构体中,实现子结构体可以访问到父类中
//UserIndex：提供给外部进行调用:admin.UserIndex
func (con IndexController) Index(c *gin.Context) {
	c.String(200, "后台首页")
	//_ :不接受err
	//获取值
	username, _ := c.Get("username")
	//get返回的是控接口类型，需要转换为string类型
	v, ok := username.(string)
	if ok == true {
		//成功：
		c.String(200, "用户列表："+v)
	} else {
		c.String(200, "用户列表：获取用户失败")
	}

	fmt.Println(username)
}
