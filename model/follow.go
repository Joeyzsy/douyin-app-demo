package model

type Follow struct {
	//gorm.Model
	Id       int64  `json:"id"`
	AuthorId string `json:"author_id"`
	FanId    string `json:"fan_id"`
}

// TableName 指定表名
func (Follow) TableName() string {
	return "fans"
}
