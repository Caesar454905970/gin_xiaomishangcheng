package main

import (
	"14_gin_demo14/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//实例化结构体
		user := models.User{
			Username: "成强",
			Age:      12,
			Email:    "454905970@qq.com",
			AddTime:  11111111111,
		}
		//创建新增(传入地址)
		models.DB.Create(&user)
		fmt.Println(user)
		c.String(200, "增加用户:%v", user)

	})
	r.GET("/news", func(c *gin.Context) {
		//保存所有字段
		user := models.User{Id: 6}
		models.DB.Find(&user)
		//更新数据
		user.Username = "你好"
		user.Email = "jinmianyiliao.com"
		models.DB.Save(user)

		fmt.Println(user)
		c.String(200, "修改用户")

	})
	r.GET("/news1", func(c *gin.Context) {
		//更新数据
		user := models.User{}
		models.DB.Model(&user).Where("id =?", 6).Update("username", "成强")
		c.String(200, "修改用户")
	})
	r.GET("/news2", func(c *gin.Context) {
		//删除数据
		user := models.User{}
		models.DB.Model(&user).Where("id =?", 6).Delete("username", "成强")
		c.String(200, "修改用户")
	})
	r.Run()
}
