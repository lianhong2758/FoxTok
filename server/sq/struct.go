package sq

import (
	"time"
)

type User struct {
	ID            int64 `gorm:"primarykey"`
	Name          string
	Follower      int64  //关注
	Fans          int64  //粉丝
	Avatar        string //头像
	Background    string //背景
	Signature     string //简介
	Favorited     int64  //被赞数
	WorkNum       int64  //作品数
	FavoriteCount int64  //点赞数
}

type PassWord struct {
	ID       int64 `gorm:"primarykey"`
	Name     string
	PassWord string
}

type Video struct {
	ID        int64 `gorm:"primarykey"`
	UserID    int64
	VideoPath string
	CoverPath string
	Favorite  int64 //点赞
	Comment   int64 //评论
	Title     string
	CreatedAt time.Time
}

type IsFavorite struct {
	ID         int64 `gorm:"primarykey"`
	UserID     int64
	VideoID    int64
	IsFavorite bool
}

type Comment struct {
	ID      int64 `gorm:"primarykey"`
	VideoID int64
	UserID  int64
	Content string
	Time    string // 评论发布日期，格式 mm-dd
}
type Message struct {
	ID      int64 `gorm:"primarykey"`
	To      int64
	From    int64
	Content string
	Time    string
}

type Follower struct {
	ID        int64 `gorm:"primarykey"`
	UserID    int64
	SubUserID int64
}
