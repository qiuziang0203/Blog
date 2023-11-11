package mysql

import (
	"Blog/entity"
	"crypto/md5"
	"encoding/hex"
)

func Register(user entity.User) (entity.User, error) {
	str := user.Password
	data := []byte(str)
	md5New := md5.New()
	md5New.Write(data)
	// hex转字符串
	md5String := hex.EncodeToString(md5New.Sum(nil))
	user.Password = md5String
	res := DB.Create(&user)
	count := entity.UserCount{}
	count.UserID = user.ID
	DB.Create(&count)
	return user, res.Error
}
func PasswordLogin(username string, password string) (entity.User, error) {
	str := password
	data := []byte(str)
	md5New := md5.New()
	md5New.Write(data)
	// hex转字符串
	md5String := hex.EncodeToString(md5New.Sum(nil))
	user := entity.User{}
	res := DB.Where("username = ? and password = ?", username, md5String).First(&user)
	return user, res.Error
}
