package video

import (
	"time"

	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/global"
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
)

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
		Id:        int(videoID),
		Title:     title,
		Play_url:  videoName,
		Cover_url: coverName,
		//FavoriteCount : 0,
		//CommentCount : 0,
		User_id:      int(userID),
		Created_time: time.Now(),
		Updated_time: time.Now(),
	}

	err := global.DB.Create(&videos).Error
	if err != nil {
		return err
	}
	return err
}
