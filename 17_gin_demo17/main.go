package main

import (
	"17_gin_demo17/models"
	"17_gin_demo17/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	//加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	r.Static("/static", "./static")

	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.DefaultRoutersInit(r)

	//演示go-ini的使用
	//config, err := ini.Load("./conf/app.ini")
	//if err != nil {
	//	//失败
	//	fmt.Printf("Fail to read file: %v", err)
	//	os.Exi t(1)
	//}
	////获取ini里面的配置
	//fmt.Println("App Mode:", config.Section("").Key("app_name").String())
	//fmt.Println("App Mode:", config.Section("mysql").Key("password").String())
	//fmt.Println("App Mode:", config.Section("redis").Key("ip").String())
	////给ini写入数据
	//config.Section("").Key("app_name").SetValue("成强")
	//config.Section("").Key("admin_path").SetValue("/admin")
	//config.SaveTo("./conf/app.ini")
	r.Run()
}
