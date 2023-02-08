package db

import "github.com/Joeyzsy/douyin-app-demo/model"

// GetVideoListById get videl list by userid
func GetVideoListById(userID int64) ([]model.Video, error) {
	var videoList []model.Video

	DB.Where("user_id = ?", userID).Find(&videoList)

	return videoList, nil
}

func GetVideosList() (*[]model.Video, error) {
	res := make([]model.Video, 0)
	if err := DB.Where("").Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func PublishVideo(videos *model.Video) error {
	err := DB.Create(&videos).Error
	if err != nil {
		return err
	}
	return err
}
