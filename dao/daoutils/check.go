package daoutils

import (
	"Blog/dao/mysql"
	"Blog/entity"
)

func IsLike(userid int64, blogid int64) bool {
	likes, _ := mysql.GetUserLike(userid)
	for i := 0; i < len(likes); i++ {
		if int(blogid) == likes[i] {
			return true
		}
	}
	return false
}
func IsFavor(userid int64, blogid int64) bool {
	likes, _ := mysql.GetUserFavor(userid)
	for i := 0; i < len(likes); i++ {
		if int(blogid) == likes[i] {
			return true
		}
	}
	return false
}

func GetBlogUser(blogid int64) (uint, error) {
	blog := entity.Blog{}
	err := mysql.DB.Where("id = ?", blogid).First(&blog).Error
	return blog.UserID, err
}
