package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
// publish 投稿接口
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	//判断用户是否存在
	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	/*videoID, err := global.ID_GENERATOR.NextID()
	if err != nil {
		// 无法生成ID
		c.JSON(http.StatusInternalServerError, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}*/
	videoID := 12

	filename := filepath.Base(data.Filename)

	user := usersLoginInfo[token]

	finalName := fmt.Sprintf("%d_%s", user.Id, filename)

	saveFile := filepath.Join("./public/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	err = service.PublishVideo(uint64(user.Id), uint64(videoID), saveFile, saveFile, filename)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
