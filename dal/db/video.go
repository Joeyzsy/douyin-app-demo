package db

import "github.com/Joeyzsy/douyin-app-demo/model"

// GetVideoListById get videl list by userid
func GetVideoListById(userID int64) ([]model.Video, error) {
	var videoList []model.Video

	DB.Where("user_id = ?", userID).Find(&videoList)

	return videoList, nil
}
