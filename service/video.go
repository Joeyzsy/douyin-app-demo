package service

import (
	"time"

	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
)

func GetvideoAndAuthor(videos *[]model.Video, authors *[]model.Users, LatestTime int64, MaxNumVideo int) (int, error) {
	//video

	var numVideos int
	global.DB.Limit(MaxNumVideo).Find(&videos)
	numVideos = len(*videos)
	return numVideos, nil
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
