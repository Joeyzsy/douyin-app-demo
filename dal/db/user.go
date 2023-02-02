package db

import (
	"github.com/Joeyzsy/douyin-app-demo/pkg/constants"
)

type User struct {
	//gorm.Model
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	UserId        int64  `json:"user_id"`
	Token         string `json:"token"`
	FollowCount   int64  `json:"follow-count"`
	FollowerCount int64  `json:"follower-count"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
func GetUserById(userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}
	if err := DB.Where("user_name in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(users []*User) (err error) {
	return DB.Create(users).Error
}

// QueryUser query list of user info
func GetUserByName(userName []string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name in ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
