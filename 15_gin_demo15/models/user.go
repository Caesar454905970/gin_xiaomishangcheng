package models

//定义 user 模型：
type User struct { // 默认表名是 `users`
	//首字母大写，增加外部访问
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

//表示把 User 结构体默认操作的表改为 user 表
func (User) TableName() string {
	return "user"
}
