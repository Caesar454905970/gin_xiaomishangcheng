package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   int
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
}

//表示配置操作数据库的表名称
func (Manager) TableName() string {
	return "manager"
}
