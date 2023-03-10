package controller

import (
	"fmt"
	"github.com/Joeyzsy/douyin-app-demo/global"
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/service/resp"
	"github.com/Joeyzsy/douyin-app-demo/service/video"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []resp.VideoResp `json:"video_list,omitempty"`
	NextTime  int64            `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	fmt.Print("feed start!\n")
	var videoList []model.Video
	var authorList []model.Users

	var CurrentTimeInt = time.Now().UnixMilli()
	var CurrentTime = strconv.FormatInt(CurrentTimeInt, 10)
	var LatestTimeStr = c.DefaultQuery("latest_time", CurrentTime)

	LatestTime, err := strconv.ParseInt(LatestTimeStr, 10, 64)
	if err != nil {
		// 无法解析latest_time
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "parameter latest_time is wrong"})
		return
	}

	response := video.GetvideoAndAuthor(&videoList, &authorList, LatestTime, global.FEED_NUM)
	if response.NumVideos == 0 {
		// 没有满足条件的视频 使用当前时间再获取一遍
		response := video.GetvideoAndAuthor(&videoList, &authorList, LatestTime, global.FEED_NUM)
		if response.NumVideos == 0 {
			// 后端没有视频了
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: 0},
				VideoList: nil,
				NextTime:  CurrentTimeInt, // 没有视频可刷时返回当前时间
			})
			return
		}
	}

	videoJsonList := make([]resp.VideoResp, 0, response.NumVideos)
	var (
		videoJson resp.VideoResp
		//author         response.Users
		authorJson     resp.User
		isFavoriteList []bool
		//isFollowList   []bool
		isLogged = false // 用户是否传入了合法有效的token（是否登录）
	)
	// 未登录时默认为未关注未点赞
	var isFavorite = false
	//var isFollow = false
	isLogged = false //此时默认未登录

	for i, video := range videoList {

		if isLogged {
			// 当用户登录时，判断是否关注当前作者
			//isFollow = isFollowList[i]
			isFavorite = isFavoriteList[i]
		}

		// 填充JSON返回值
		/*
			author = authorList[i]
			authorJson.Id = int64(author.Id)
			authorJson.Name = author.Name
			authorJson.FollowCount = author.FollowCount
			authorJson.FollowerCount = author.FollowerCount
			authorJson.TotalFavorited = author.TotalFavorited
			authorJson.FavoriteCount = author.FavoriteCount
			authorJson.IsFollow = isFollow
		*/
		videoJson.Id = int64(video.Id)
		videoJson.Author = authorJson
		videoJson.PlayUrl = video.PlayUrl
		videoJson.CoverUrl = video.CoverUrl
		videoJson.FavoriteCount = int64(video.FavoriteCount)
		videoJson.CommentCount = int64(video.CommentCount)
		//videoJson. = video.Title
		videoJson.IsFavorite = isFavorite

		videoJsonList = append(videoJsonList, videoJson)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoJsonList,
		NextTime:  time.Now().Unix(),
	})
}
