package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
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
		file, err := c.FormFile("face")
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

	r.POST("/admin/user/doeidt", func(c *gin.Context) {
		username := c.PostForm("username")
		file1, err1 := c.FormFile("face1")
		dst1 := path.Join("./static/upload", file1.Filename)
		if err1 == nil {
			//成功
			//存储文件到本地
			c.SaveUploadedFile(file1, dst1)
		}
		file2, err2 := c.FormFile("face2")
		dst2 := path.Join("./static/upload", file2.Filename)
		if err2 == nil {
			//成功
			//存储文件到本地
			c.SaveUploadedFile(file2, dst2)
		}
		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
			"dst1":     dst1,
			"dst2":     dst2,
		})
	})
	r.POST("/admin/user/doeidtnames", func(c *gin.Context) {
		username := c.PostForm("username")
		form, _ := c.MultipartForm()
		files := form.File["face[]"]
		for _, file := range files {
			// 上传文件至指定目录
			dst := path.Join("./static/upload", file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
		})
	})

	r.Run(":8080")

}
