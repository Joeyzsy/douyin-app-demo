package controller

import (
	"fmt"
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
	"github.com/Joeyzsy/douyin-app-demo/service/follow"
	"github.com/Joeyzsy/douyin-app-demo/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	service := user.UserServiceImpl{}
	resp := service.RegisterUser(username, password)

	if resp.ReturnErr == errno.UserAlreadyExistErr {
		// search for user_name to see if exists
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else if resp.ReturnErr == errno.Success {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   resp.Userid,
			Token:    resp.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Service error"},
		})
	}
}

func Login(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")

	username := "plus"
	password := "123456"

	fmt.Println(username, " ", password)

	service := user.UserServiceImpl{}
	resp := service.LoginUser(username, password)
	if resp.ReturnErr == errno.Success {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   resp.Userid,
			Token:    resp.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: resp.ReturnErr.ErrMsg}})
	}
}

func UserInfo(c *gin.Context) {
	userService := user.UserServiceImpl{}
	followService := follow.FollowServiceImpl{}

	// 获取指定用户的 ID
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: int32(errno.ParamErr.ErrCode), StatusMsg: errno.ParamErr.ErrMsg})
		return
	}

	// 获取指定用户的信息
	userModelResp := userService.GetUserInfo(int64(userID))
	if userModelResp.ReturnErr != errno.Success {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: errno.ParamErr.ErrMsg})
		return
	}

	// 获取当前用户的 ID
	curUserID := c.GetUint64("UserID")

	// 判断当前用户是否关注指定用户
	followStatusResp := followService.GetFollowStatus(int64(curUserID), int64(userID))
	if followStatusResp.ReturnErr != errno.Success {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: errno.ParamErr.ErrMsg})
		return
	}

	userModelResp.User.IsFollow = followStatusResp.IsFollow

	var resp = UserResponse{
		Response: Response{StatusCode: 0, StatusMsg: "OK"},
	}
	resp.User = userModelResp.User
	c.JSON(http.StatusOK, resp)
}
