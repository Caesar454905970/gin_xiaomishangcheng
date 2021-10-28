package admin

import (
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	BaseController
}

func (con StudentController) Index(c *gin.Context) {
	//1、获取所有的学生信息
	//studentList := []models.Student{}
	//models.DB.Find(&studentList)
	//c.JSON(200, gin.H{
	//	"result": studentList,
	//})

	//2、获取所有的课程信息

	// lessonList := []models.Lesson{}
	// models.DB.Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })

	//3、查询学生信息的时候 展示学生选修的课程
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Find(&studentList)
	// c.JSON(200, gin.H{
	// 	"result": studentList,
	// })

	// 4、查询张三 以及张三选修了哪些课程

	//studentList := []models.Student{}
	//models.DB.Preload("Lesson").Where("name = ?", "张三").Find(&studentList)
	//c.JSON(200, gin.H{
	//	"result": studentList,
	//})

	// 5、查询课程被哪些学生选修了

	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Find(&lessonList)
	//c.JSON(200, gin.H{
	//	"result": lessonList,
	//})

	// 6、查询计算机网络被哪些学生选修了
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Where("name=?", "计算机网络").Find(&lessonList)
	//c.JSON(200, gin.H{
	//	"result": lessonList,
	//})

	// 7、查询数据指定条件
	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student").Offset(1).Limit(2).Order("id desc").Find(&lessonList)
	//c.JSON(200, gin.H{
	//	"result": lessonList,
	//})

	// 7、张三被开除了  查询课程被哪些学生选修的时候去掉张三

	// lessonList := []models.Lesson{}
	// models.DB.Preload("Student", "id != ?", 1).Find(&lessonList)
	// c.JSON(200, gin.H{
	// 	"result": lessonList,
	// })

	// 8、张三 李四都被开除了  查询课程被哪些学生选修的时候去掉张三和李四

	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student", "id not in (1,2)").Find(&lessonList)
	//c.JSON(200, gin.H{
	//	"result": lessonList,
	//})

	// 9 查看课程被哪些学生选修 要求：学生 id 倒叙输出  自定义预加载 SQL
	// https://gorm.io/zh_CN/docs/preload.html

	//lessonList := []models.Lesson{}
	//models.DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
	//	return models.DB.Where("id>2").Order("student.id DESC")
	//}).Find(&lessonList)
	//c.JSON(200, gin.H{
	//	"result": lessonList,
	//})

}
