package service

import (
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
