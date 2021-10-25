package itying

import "github.com/gin-gonic/gin"

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	//设置Cookie
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, false)
	c.String(200, "我是一个api接口")
}
func (con DefaultController) News(c *gin.Context) {

	//获取Cookie
	username, _ := c.Cookie("username")
	c.String(200, "获取的cookie是"+username)
}
