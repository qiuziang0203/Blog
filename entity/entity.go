package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Status   int64  `json:"status"` //0-普通用户 1-管理员 2-被封禁用户
	NickName string `json:"nick_name"`
	Sign     string `json:"sign"`
	Birth    string `json:"birth"`
	Sex      string `json:"sex"`
}

type UserCount struct {
	gorm.Model
	UserID        uint  `json:"user_id"`
	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	WorkCount     int64 `json:"work_count"`
	FavorCount    int64 `json:"favor_count"`
}

type Blog struct {
	gorm.Model
	UserID     uint
	Title      string
	Text       string
	Status     int // 0-草稿 1-待审核 2-已发布
	LikeNum    int
	FavorNum   int
	CommentNum int
	Hot        int
}

type Type struct {
	gorm.Model
	TypeName   string
	TypeBrief  string
	TypeAvatar string
}
type BlogType struct {
	gorm.Model
	BlogID uint
	TypeID uint
	Hot    int
}

type Like struct {
	gorm.Model
	UserID uint
	BlogID uint
}
type Favor struct {
	gorm.Model
	UserID uint
	BlogID uint
}
type Comment struct {
	gorm.Model
	UserID       uint
	BlogID       uint
	Text         string
	LikeNum      int
	ReCommentNum int
}
type ReComment struct {
	gorm.Model
	UserID    uint
	CommentID uint
	Text      string
	LikeNum   int
}
type Follow struct {
	gorm.Model
	FollowID   uint
	FollowerID uint
}
type CommentLike struct {
	gorm.Model
	UserID    uint
	CommentID uint
}
type Announcement struct {
	gorm.Model
	Title string
	Text  string
}
