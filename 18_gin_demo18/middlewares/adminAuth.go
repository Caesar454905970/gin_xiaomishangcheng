package middlewares

import (
	"18_gin_demo18/models"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

//后端管理页面进行session读取，进行权限判断
func InitAdminAuthMiddleware(c *gin.Context) {
	//没有登录的用户，不能进入用户管理中心
	fmt.Println("进行鉴权中间件:InitAdminAuthMiddleware")
	//进行权限判断 没有登录的用户 不能进入后台管理中心
	//1、获取Url访问的地址  /admin/captcha
	//2、获取Session里面保存的用户信息
	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行
	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	//  1、获取Url访问的地址   /admin/captcha?t=0.8706946438889653
	//去掉get传值
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	//fmt.Println("1、获取Url访问的地址:", pathname)

	//2、获取Session里面保存的用户信息
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo") //获取的session中保存的用户登录信息
	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string) //判断
	if ok {
		//是字符串
		var userinfoStruct []models.Manager
		//把字符转换成结构体(Manager)对应的切片
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		//fmt.Println(userinfoStruct)
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				//没有登录信息
				//排除用户当前在用户登录的路由界面
				c.JSON(200, gin.H{
					"code":           400,
					"msg":            "session不存在，没有登录信息，告诉前端返回到登录页面RedirectUrl",
					"userinfoStruct": "",
					"RedirectUrl":    "/admin/login",
				})
			}
		} else {
			//有登录信息
			c.JSON(200, gin.H{
				//"username":userinfoStruct[0].Username,
				"code":           200,
				"msg":            "session，存在登录信息，放行通过",
				"userinfoStruct": userinfoStruct,
				"currenturl":     pathname,
			})
		}
	} else {
		//用户没有登录，告诉前端返回到登录页面
		//排除用户当前在用户登录的路由界面
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.JSON(200, gin.H{
				"code":           400,
				"msg":            "session不存在，没有登录信息，告诉前端返回到登录页面RedirectUrl",
				"userinfoStruct": "",
				"RedirectUrl":    "/admin/login",
			})

		}

	}

}
