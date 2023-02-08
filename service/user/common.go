package user

import (
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
)

type UserRegisterResponse struct {
	Userid    int64
	Token     string
	ReturnErr errno.ErrNo
}

type UserInfoResponse struct {
	User      model.User
	ReturnErr errno.ErrNo
}
