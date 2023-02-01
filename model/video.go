package model

import (
	"time"
)

type Video struct {
	Id             int       `json:"id" gorm:"id"`
	User_id        int       `json:"user_id" gorm:"user_id"`
	Play_url       string    `json:"paly_url" gorm:"paly_url"`
	Cover_url      string    `json:"cover_url" gorm:"cover_url"`
	Title          string    `json:"title" gorm:"title"`
	Favorite_count int       `json:"favorite_count" gorm:"favorite_count"`
	Comment_count  int       `json:"comment_count" gorm:"comment_count"`
	Created_time   time.Time `json:"created_time" gorm:"created_time"`
	Updated_time   time.Time `json:"updated_time" gorm:"updated_time"`
}
