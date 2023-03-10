package controller

import (
	"fmt"
	"github.com/Joeyzsy/douyin-app-demo/model"
	"github.com/Joeyzsy/douyin-app-demo/pkg/errno"
	"github.com/Joeyzsy/douyin-app-demo/service/follow"
	"github.com/Joeyzsy/douyin-app-demo/service/resp"
	"github.com/Joeyzsy/douyin-app-demo/service/user"
	"github.com/gin-gonic/gin"
	"log"
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
		//IsFollow:      true,
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
	User resp.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	service := user.UserServiceImpl{}
	resp := service.RegisterUser(username, password)

	if resp.ReturnErr == errno.UserAlreadyExistErr {
		// search for user_name to see if exists
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: resp.ReturnErr.ErrMsg},
		})
	} else if resp.ReturnErr == errno.Success {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   resp.Userid,
			Token:    resp.Token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: resp.ReturnErr.ErrMsg},
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

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
			Response: Response{StatusCode: 1, StatusMsg: "Service error"}})
	}
}

func UserInfo(c *gin.Context) {
	userService := user.UserServiceImpl{}
	followService := follow.FollowServiceImpl{}

	// ????????????????????? ID
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: int32(errno.ParamErr.ErrCode), StatusMsg: errno.ParamErr.ErrMsg})
		return
	}

	// ??????User Service, ???????????????????????????
	userModelResp := userService.GetUserInfo(int64(userID))
	if userModelResp.ReturnErr != errno.Success {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: errno.ParamErr.ErrMsg})
		return
	}

	// ????????????????????? ID
	curUserID := c.GetUint64("UserID")

	// ??????Follow Service, ??????????????????????????????????????????
	followStatusResp := followService.GetFollowStatus(int64(curUserID), int64(userID))
	if followStatusResp.ReturnErr != errno.Success {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: errno.ParamErr.ErrMsg})
		return
	}

	userModelResp.User.IsFollow = followStatusResp.IsFollow

	log.Println("In userinfo controller---user: ", userModelResp.User)
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0, StatusMsg: "OK"},
		User:     userModelResp.User,
	})

}
