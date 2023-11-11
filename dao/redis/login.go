package redis

import (
	"Blog/entity"
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// 静态信息存入redis
func Register(user entity.User) error {
	str := user.Password
	d := []byte(str)
	md5New := md5.New()
	md5New.Write(d)
	// hex转字符串
	md5String := hex.EncodeToString(md5New.Sum(nil))

	ctx := context.Background()
	data := make(map[string]interface{})
	data["id"] = user.ID
	data["username"] = user.Username
	data["password"] = md5String
	data["nickname"] = user.NickName
	data["status"] = user.Status
	data["email"] = user.Email
	data["avatar"] = user.Avatar

	err := Rdb.HMSet(ctx, "user_info::user_id"+strconv.FormatInt(int64(user.ID), 10), data).Err()
	Rdb.Expire(ctx, "user_info::user_id"+strconv.FormatInt(int64(user.ID), 10), time.Hour*24*7)
	return err
}

// 登录时token存入redis
func SetToken(id int64, token string) {
	ctx := context.Background()
	Rdb.Set(ctx, "user_token::user_id"+strconv.FormatInt(id, 10), token, time.Hour*24*7)
}

// 获取redis中的token
func GetToken(id int64) (string, error) {
	ctx := context.Background()
	token, err := Rdb.Get(ctx, "user_token::user_id"+strconv.FormatInt(id, 10)).Result()
	return token, err
}

// 退出登录，将redis中的token删除
func Logout(id int64) error {
	ctx := context.Background()
	err := Rdb.Del(ctx, "user_token::user_id"+strconv.FormatInt(id, 10)).Err()
	return err
}
