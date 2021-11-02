package admin

import (
	"18_gin_demo18/models"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	//验证Md5
	fmt.Println(models.Md5("123456"))
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {
	//获取前端传过来验证码的CaptchaId和verifyValue
	captchId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(models.Md5("123456"))
	//1、验证验证码是否正确（防止别人进行工具)
	//排除前端提交的空请求(字符串前后空格)strings.Trim(str," ")
	if strings.Trim(captchId, " ") == "" || strings.Trim(verifyValue, " ") == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":              400,
			"msg":               "验证码验证失败,前端提交的空请求(字符串前后空格)",
			"verifyValueResult": false,
		})
		return
	}
	//调用验证验证码的方法
	if flag := models.VerifyCaptcha(captchId, verifyValue); flag == true {
		//验证通过：再去执行登录操作

		//2、查询数据库 判断用户以及密码是否存在
		userinfoList := []models.Manager{}
		//对前端传过来的密码进行Md5加密
		password = models.Md5(password)
		//把查询结果=所有字段保存到userinfo(传入地址)

		models.DB.Where("username = ? AND password = ?", username, password).Find(&userinfoList)

		fmt.Println("查询数据库成功")
		//fmt.Println(username)
		//fmt.Println(password)
		//if username =="admin" && password =="123456" {
		if len(userinfoList) > 0 {

			//3、执行登录 保存用户信息(cookies:客户端 sessions:服务器端) 执行跳转

			//创建sesson
			session := sessions.Default(c)

			//保存用户信息，sessions:服务器端
			//session.Set不能直接保存切片,把结构体转换成json字符串
			userinfoSlice, _ := json.Marshal(userinfoList)
			session.Set("userinfo", string(userinfoSlice))
			session.Save()

			//提示登录成功信息
			c.JSON(http.StatusOK, gin.H{
				"code":              200,
				"msg":               "验证码验证成功，登录成功",
				"verifyValueResult": "true",
				"login":             "true",
			})
		} else {
			//登录失败
			c.JSON(http.StatusOK, gin.H{
				"code":              400,
				"msg":               "验证码验证成功,用户名密码错误",
				"verifyValueResult": "true",
				"login":             "false",
			})
		}

		fmt.Println(flag)

	} else {
		//验证失败
		c.JSON(http.StatusOK, gin.H{
			"code":              400,
			"msg":               "验证码验证失败",
			"verifyValueResult": "false",
			"login":             "false",
		})
	}

}
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		//失败
		fmt.Println(err)
	}
	type code struct {
		BaseController
	}
	//结构体返回，灵活使用tag来对结构体字段做定制化操作
	//type msg struct {
	//	Code     int `json:"code"`
	//	Msg    string `json:"msg"`
	//	CaptchaId string `json:"captchaId"`
	//	CaptchaImage     string `json:"captchaImage"`
	//}
	//msg1 :=msg{
	//	Code:         200,
	//	Msg:          "获取验证码验证成功",
	//	CaptchaId:    id,
	//	CaptchaImage: b64s, //验证码的url地址
	//}
	c.JSON(200, gin.H{
		"code":         200,
		"msg":          "获取验证码验证成功",
		"CaptchaId":    id,
		"CaptchaImage": b64s, //验证码的url地址
	})
}

func (con LoginController) LoginOut(c *gin.Context) {
	//销毁session
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "退出session成功，删除session",
	})
}
