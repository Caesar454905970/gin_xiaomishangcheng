package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type userInfo struct {
	//首字母大写表示公有的
	Username string `json:"username" form:"username" `
	Password string `json:"password" form:"password" `
}
type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// UnixToTime 时间戳转换成日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}
func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)

	return str1 + "-----" + str2
}

func main() {
	r := gin.Default()

	//配置自定义模板函数
	r.SetFuncMap(template.FuncMap{
		//注册模板函数
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//配置模板文件的路径,放在路由配置的前面:加载templates下的所有文件
	r.LoadHTMLGlob("templates/**/*")

	//静态资源，第一个参数是路由，第二个是本地的目录
	r.Static("/static", "./static")
	//Get 请求传值
	//请求方式：http://localhost:8080/?username=1&age=12&password=123
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		//DefaultQuery：有返回值赋值，不存在给默认值1
		page := c.DefaultQuery("page", "1")
		password := c.Query("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
			"password": password,
		})
	})
	//Get请求传值 id
	//http://localhost:8080/article?id=12
	r.GET("/article", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"mag": "新闻详情",
			"id":  id,
		})
	})

	//post演示
	//http://localhost:8080/article?id=12
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "itying/user.html", gin.H{})
	})
	r.POST("/doAddUser", func(c *gin.Context) {
		//获取表单post过来的数据:一个一个获取
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	//获取 GET POST 传递的数据绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		//实例化一个结构体
		user := &userInfo{}
		//ShouldBind:必须传入地址
		if err := c.ShouldBind(&user); err == nil {
			//成功
			fmt.Printf("%#v\n", user) //打印结构体数据
			c.JSON(http.StatusOK, user)
		} else {
			//失败
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	//获取 Post Xml 数据
	r.POST("/xml", func(c *gin.Context) {
		//实例化一个结构体
		article := &Article{}
		xmlSliceData, _ := c.GetRawData() // 从 c.Request.Body 读取xml请求数据
		if err := xml.Unmarshal(xmlSliceData, article); err == nil {
			//成功
			c.JSON(http.StatusOK, article)
		} else {
			//失败
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	//动态路由传值
	//localhost:8080/list/123
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.JSON(http.StatusOK, gin.H{
			"cid": cid,
		})
	})
	r.Run()

}
