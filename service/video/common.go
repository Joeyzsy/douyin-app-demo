package video

import "github.com/Joeyzsy/douyin-app-demo/pkg/errno"

type GetVideoResponse struct {
	NumVideos int
	ReturnErr errno.ErrNo
}
