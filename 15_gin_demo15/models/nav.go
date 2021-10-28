package models

type Nav struct {
	//`json:"id"`:返回的参数大写变小写
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	Status int    `json:"status"`
	Sort   int    `json:"sort"`
}

func (Nav) TableName() string {
	return "nav"
}
