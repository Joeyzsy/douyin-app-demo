package user

import (
	"github.com/Joeyzsy/douyin-app-demo/dal/db"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

func (s *UserServiceImpl) LoginUser(name string, pwd string) (resp UserRegisterResponse) {
	resp = UserRegisterResponse{0, "", errno.Success}

	if len(name) == 0 || len(pwd) == 0 {
		resp.ReturnErr = errno.ParamErr
		return resp
	}

	users, err := db.GetUserByName([]string{name})
	if err != nil || (len(users) != 0 && len(users) != 1) {
		resp.ReturnErr = errno.ServiceErr
		return resp
	}
	if len(users) == 0 {
		resp.ReturnErr = errno.UserNonExistErr
		return resp
	}
	if users[0].Password != pwd {
		resp.ReturnErr = errno.PwdWrongErr
		return resp
	}
	resp.Userid = users[0].UserId
	resp.Token = users[0].Token
	return resp
}

func (s *UserServiceImpl) RegisterUser(name string, pwd string) (resp UserRegisterResponse) {

	resp = UserRegisterResponse{0, "", errno.Success}

	if len(name) == 0 || len(pwd) == 0 {
		resp.ReturnErr = errno.ParamErr
		return resp
	}

	users, err := db.GetUserByName([]string{name})
	if err != nil || (len(users) != 0 && len(users) != 1) {
		resp.ReturnErr = errno.ServiceErr
		return resp
	}
	if len(users) != 0 {
		resp.ReturnErr = errno.UserAlreadyExistErr
		return resp
	}

	userid := GenerateUserId()
	token := GenerateToken(name, pwd)
	user := &db.User{UserName: name, Password: pwd, UserId: userid, Token: token}
	err = db.CreateUser([]*db.User{user})
	if err != nil {
		resp.ReturnErr = errno.ServiceErr
		return resp
	}

	resp.Userid = userid
	resp.Token = token

	return resp
}

func GenerateUserId() (random int64) {
	return time.Now().Unix()
}

func GenerateToken(username string, pwd string) (random string) {
	return username + pwd
}