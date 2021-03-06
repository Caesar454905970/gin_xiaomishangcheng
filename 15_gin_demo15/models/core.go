package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	//读取.ini里面的数据库配置
	config, err := ini.Load("./conf/app.ini")
	if err != nil {
		//失败
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	ip := config.Section("mysql").Key("ip").String()
	port := config.Section("mysql").Key("port").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	database := config.Section("mysql").Key("database").String()

	fmt.Println("App Mode:", config.Section("mysql").Key("password").String())
	fmt.Println("App Mode:", config.Section("redis").Key("ip").String())

	//dsn := "gin:gin@tcp(111.229.91.20:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		ip,
		port,
		database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true, //打印sql
		SkipDefaultTransaction: true, //禁用mysql事务
	})
	// DB.Debug()
	if err != nil {
		fmt.Println(err)
	}
}
