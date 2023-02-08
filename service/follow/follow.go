package follow

import (
	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

func (s *FollowServiceImpl) GetFollowStatus(curUserId int64, userId int64) (resp FollowStatusResponse) {
	resp = FollowStatusResponse{false, errno.Success}

	isFollow, err := db.GetFollowStatus(curUserId, userId)
	if err != nil {
		resp.ReturnErr = errno.ServiceErr
		return resp
	}

	resp.IsFollow = isFollow

	return resp
}
