package video

import (
	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/service/user"
)

type VideoServiceImpl struct{}

func (s *VideoServiceImpl) GetPublishedVideosByUserId(userId int64) ([]model.VideoResp, error) {
	videoList, err := db.GetVideoListById(userId)

	res := make([]model.VideoResp, len(videoList))

	userservice := user.UserServiceImpl{}

	userModelResp := userservice.GetUserInfo(userId)
	userEntity := userModelResp.User

	var i int
	for i = 0; i < len(videoList); i++ {
		var video model.VideoResp

		video.Id = int64(videoList[i].Id)
		video.Author = userEntity
		video.CommentCount = int64(videoList[i].CommentCount)
		video.CoverUrl = videoList[i].CoverUrl
		video.FavoriteCount = int64(videoList[i].FavoriteCount)
		video.Title = videoList[i].Title
		video.PlayUrl = videoList[i].PlayUrl

		res[i] = video
	}

	return res, err
}
