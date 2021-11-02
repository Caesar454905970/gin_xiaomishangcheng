package admin

import (
	"17_gin_demo17/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})

}
func (con LoginController) DoLogin(c *gin.Context) {
	//获取前端传过来的CaptchaId和verifyValue
	captchId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	//排除前端提交的空请求(字符串前后空格)strings.Trim(str," ")
	if strings.Trim(captchId, " ") == "" || strings.Trim(verifyValue, " ") == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":              400,
			"msg":               "验证码验证失败",
			"verifyValueResult": false,
		})
		return
	}
	//调用验证验证码的方法
	if flag := models.VerifyCaptcha(captchId, verifyValue); flag == true {
		//验证通过
		fmt.Println(flag)
		c.JSON(http.StatusOK, gin.H{
			"code":              200,
			"msg":               "验证码验证成功",
			"verifyValueResult": "true",
		})
	} else {
		//验证失败
		c.JSON(http.StatusOK, gin.H{
			"code":              400,
			"msg":               "验证码验证失败",
			"verifyValueResult": false,
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
