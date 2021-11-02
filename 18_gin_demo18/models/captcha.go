package models

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 设置自带的 store（可以配置成redis）
//var store = base64Captcha.DefaultMemStore

//配置RedisStore RedisStore实现base64Captcha.Store的接口
var store base64Captcha.Store = RedisStore{}

//获取验证码
func MakeCaptcha() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	//配置验证码的参数
	//driverString := base64Captcha.DriverChinese{//中文验证码
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          2, //验证码长度
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		//Source: "生成的就是中文验证码这里面的文字是配置文字源的",
		BgColor: &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:   []string{"wqy-microhei.ttc"},
	}
	//ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()
	//创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	//Generate 生成随机 id、base64 图像字符串
	id, b64s, err = captcha.Generate()
	return id, b64s, err

}

//验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	fmt.Println(id, VerifyValue)
	if store.Verify(id, VerifyValue, true) {
		//验证成功
		return true
	} else {
		//验证失败
		return false
	}
}
