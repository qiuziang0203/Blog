package utils

import "fmt"

var (
	username = "root"         // 账号
	password = "root"         //密码
	host     = "39.101.1.119" // 数据库地址
	port     = 3306           //端口号
	Dbname   = "blog"         //数据库名
	timeout  = "10s"          // 连接超时时间

	DSN = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	TOKEN_EXPIRE_TIME = 60 * 60 * 24
)
