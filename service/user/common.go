package user

import (
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
	"github.com/Joeyzsy/douyin-app-demo/service/resp"
)

type UserRegisterResponse struct {
	Userid    int64
	Token     string
	ReturnErr errno.ErrNo
}

type UserInfoResponse struct {
	User      resp.User
	ReturnErr errno.ErrNo
}
