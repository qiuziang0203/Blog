package dao

import (
	"Blog/dao"
	"Blog/dao/mysql"
	"Blog/dao/redis"
	"Blog/entity"
	"fmt"
	"testing"
)

func Test_Creat(t *testing.T) {
	mysql.Init()
	mysql.DB.AutoMigrate(&entity.Announcement{})
}

func Test_RedisLogin(t *testing.T) {
	dao.Init()
	user := entity.User{Username: "testuser4", Password: "qza123456"}
	user.ID = 8
}

func Test_Limit(t *testing.T) {
	dao.Init()
	blogs, count, err := mysql.GetBlogByType("后端", "id desc", 1)
	fmt.Println(blogs)
	fmt.Println(count)
	fmt.Println(err)
}

func Test_HSGet(t *testing.T) {
	dao.Init()
	redis.LikeBlog(4, 4, 4)
}
