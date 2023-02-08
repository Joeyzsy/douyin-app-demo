package model

type Users struct {
	Id             int    `json:"id" gorm:"id"`
	User_id        int    `json:"user_id" gorm:"user_id""`
	User_name      string `json:"user_name" gorm:"user_name"`
	Password       string `json:"password" gorm:"password"`
	Token          string `json:"token" gorm:"token"`
	Follow_count   int    `json:"follow_count" gorm:"follow_count"`
	Follower_count int    `json:"follower_count" gorm:"follower_count"`
}
