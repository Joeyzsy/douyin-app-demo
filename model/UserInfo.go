package model

type UserInfo struct {
	Id            int64       `json:"id" gorm:"id,omitempty"`
	Name          string      `json:"name" gorm:"user_name,omitempty"`
	Password      string      `json:"password" gorm:"password,omitempty"`
	FollowCount   int64       `json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount int64       `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow      bool        `json:"is_follow" gorm:"is_follow,omitempty"`
	Videos        []*Video    `json:"-"`                                   //用户与投稿视频的一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:fans;"`            //用户之间的多对多
	FavorVideos   []*Video    `json:"-" gorm:"many2many:my_liked_videos;"` //用户与点赞视频之间的多对多

}
