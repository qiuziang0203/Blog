package dao

import (
	"Blog/dao/mysql"
	"Blog/dao/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
