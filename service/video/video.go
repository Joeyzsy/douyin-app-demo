package video

import (
	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/model"
)

type VideoServiceImpl struct{}

func (s *VideoServiceImpl) GetPublishedVideosByUserId(userId int64) ([]model.Video, error) {
	videoList, err := db.GetVideoListById(userId)

	return videoList, err
}
