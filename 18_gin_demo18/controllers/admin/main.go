package admin

import "C"
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	/*	//获取userinfo 对应的session
		session := sessions.Default(c)
		userinfo := session.Get("userinfo") //获取的session中保存的用户登录信息

		//类型断言 来判断 userinfo是不是一个string
		userinfoStr, ok := userinfo.(string) //判断
		if ok {
			//是字符串
			var userinfoStruct []models.Manager
			//把字符转换成结构体(Manager)对应的切片
			json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
			fmt.Println(userinfoStruct)
			c.JSON(200, gin.H{
				//"username":userinfoStruct[0].Username,
				"userinfoStruct": userinfoStruct,
			})
		} else {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "session不存在",
			})
		}*/

	//c.HTML(http.StatusOK, "admin/main/index.html", gin.H{})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
