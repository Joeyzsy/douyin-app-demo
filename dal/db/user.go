package db

import (
	"github.com/Joeyzsy/douyin-app-demo/model"
)

// GetUserById get user info by userid
func GetUserById(userID int64) (model.User, error) {
	var res model.User

	DB.Where("id = ?", userID).Find(&res)

	return res, nil
}

// CreateUser create user info
func CreateUser(users []*model.User) (err error) {
	return DB.Create(users).Error
}

// GetUserByName QueryUser query list of user info
func GetUserByName(userName []string) ([]*model.User, error) {
	res := make([]*model.User, 0)

	if err := DB.Where("user_name in ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
