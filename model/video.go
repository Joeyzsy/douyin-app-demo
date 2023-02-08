package model

import (
	"time"
)

type Video struct {
	Id            int       `json:"id" gorm:"id"`
	UserId        int       `json:"user_id" gorm:"user_id"`
	PlayUrl       string    `json:"play_url" gorm:"play_url"`
	CoverUrl      string    `json:"cover_url" gorm:"cover_url"`
	Title         string    `json:"title" gorm:"title"`
	FavoriteCount int       `json:"favorite_count" gorm:"favorite_count"`
	CommentCount  int       `json:"comment_count" gorm:"comment_count"`
	CreatedTime   time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime   time.Time `json:"updated_time" gorm:"updated_time"`
}
