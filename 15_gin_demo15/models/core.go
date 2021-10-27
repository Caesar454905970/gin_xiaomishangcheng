package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//gorm初始化数据库
var DB *gorm.DB
var err error

func init() {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "gin:gin@tcp(111.229.91.20:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	//全局使用，定义成共有的
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//连接失败
		fmt.Println(err)
	} else {
		//成功
		fmt.Println(DB)
	}

}
