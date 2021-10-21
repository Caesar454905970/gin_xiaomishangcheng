package main

import (
	"11_gin_demo11/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

func main() {
	r := gin.Default()
	//配置模板文件的路径,放在路由配置的前面:加载templates下的所有文件
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/user/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
	})
	r.POST("/admin/user/doUpload", func(c *gin.Context) {
		username := c.PostForm("username")
		//1、获取上传的文件
		file, err := c.FormFile("face")
		if err == nil {
			//2、获取文件的后缀名 判断类型是否正确 .jpg .png .gif .jpeg
			extName := path.Ext(file.Filename)

			//允许上传的格式
			allowExtMap := map[string]bool{
				".jpg":  true,
				".png":  true,
				".gif":  true,
				".jpeg": true,
			}
			if _, ok := allowExtMap[extName]; !ok {
				c.JSON(200, gin.H{
					"msg": "上传的文件不合法",
				})
				//退出
				return
			}
		}
		//3、创建图片保存的目录  static /upload/2021/1021
		//	获取今天的日期
		day := models.GetDate()
		dir := "./static/upload/" + day
		if err := os.MkdirAll(dir, 0666); err != nil {
			//创建文件夹失败
			c.JSON(200, gin.H{
				"msg": "MkdirAll失败",
			})
		}
		//生产文件名称和文件保存的目录
		//文件名 = 获取时间戳 +上传接收到的文件名称
		fileName := strconv.FormatInt(models.GetUnix(), 10) + extName
		fmt.Printlin(extName)
		//5、执行上传
		dst := path.Join("./static/upload", file.Filename)
		if err == nil {
			//成功
			//存储文件到本地
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(200, gin.H{
			"msg": "上传成功",
		})
		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
			"dst":      dst,
		})
	})

	r.Run(":8080")

}
