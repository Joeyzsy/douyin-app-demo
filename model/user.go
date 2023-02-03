package model

type User struct {
	//gorm.Model
	Id            int64  `gorm:"primaryKey" json:"user_id"`
	Name          string `gorm:"column:user_name"  json:"user_name"`
	Password      string `json:"password"`
	Token         string `json:"token"`
	FollowCount   int64  `json:"follow-count"`
	FollowerCount int64  `json:"follower-count"`
	IsFollow      bool
}
