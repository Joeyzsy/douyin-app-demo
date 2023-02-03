package db

import "github.com/Joeyzsy/douyin-app-demo/model"

func GetFollowStatus(curUserId int64, userId int64) (bool, error) {
	var follow model.Follow

	DB.Where("author_id = ? AND fan_id = ?", userId, curUserId).Find(&follow)

	return follow != model.Follow{}, nil
}
