package video

import (
	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/global"
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
	"github.com/Joeyzsy/douyin-app-demo/service/user"
	"time"
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

func GetvideoAndAuthor(videos *[]model.Video, authors *[]model.Users, LatestTime int64, MaxNumVideo int) (resp GetVideoResponse) {

	var numVideos int
	videos, err := db.GetVideosList()
	if err != nil || (len(*videos) == 0) {
		resp.ReturnErr = errno.ServiceErr
		return resp
	}
	numVideos = len(*videos)
	resp.NumVideos = numVideos

	return resp
}
func PublishVideo(userID uint64, videoID uint64, videoName string, coverName string, title string) error {
	videos := model.Video{
		Id:       int(videoID),
		Title:    title,
		PlayUrl:  videoName,
		CoverUrl: coverName,
		//FavoriteCount : 0,
		//CommentCount : 0,
		UserId:      int(userID),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err := global.DB.Create(&videos).Error
	if err != nil {
		return err
	}
	return err
}
