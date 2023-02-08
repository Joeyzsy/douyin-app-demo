package model

type User struct {
	//gorm.Model
	Id            int64  `gorm:"-;primary_key;AUTO_INCREMENT"`
	UserId        int64  `gorm:"user_id"`
	Name          string `gorm:"column:user_name"`
	Password      string `gorm:"password"`
	Token         string `gorm:"token"`
	FollowCount   int64  `gorm:"follow_count"`
	FollowerCount int64  `gorm:"follower_count"`
	//IsFollow      bool
}
