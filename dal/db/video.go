package db

import (
	"github.com/Joeyzsy/douyin-app-demo/model"
)

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
