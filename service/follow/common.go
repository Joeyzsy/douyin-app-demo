package follow

import "github.com/Joeyzsy/douyin-app-demo/pkg/errno"

type FollowStatusResponse struct {
	IsFollow  bool
	ReturnErr errno.ErrNo
}
