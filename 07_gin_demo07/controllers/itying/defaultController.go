package itying

import "github.com/gin-gonic/gin"

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(200, "我是一个api接口")
}
func (con DefaultController) News(c *gin.Context) {
	c.String(200, "News")
}
