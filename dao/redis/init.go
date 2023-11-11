package redis

import "github.com/go-redis/redis/v8"

var Rdb *redis.Client

func Init() {
	//ctx := context.Background()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "39.101.1.119:6379",
		Password: "qza040203", // 没有密码，默认值
		DB:       0,           // 默认DB 0
	})

	// 清空数据库
	//Rdb.FlushAll(ctx)
}
