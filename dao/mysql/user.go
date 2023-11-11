package mysql

import (
	"Blog/entity"
	"errors"
)

func GetUserByID(id int64) entity.User {
	user := entity.User{}
	DB.Where("id = ?", id).First(&user)
	return user
}

func GetUserStatus(id int) (int64, error) {
	user := entity.User{}
	res := DB.Where("id = ?", id).First(&user)
	var err error
	if res.RowsAffected == 0 {
		err = errors.New("用户不存在")
	}
	return user.Status, err
}

func GetUserByKeyword(keyword string, pagenum int64) ([]entity.User, int64, error) {
	users := []entity.User{}
	users1 := []entity.User{}
	count := DB.Where("nick_name like ? and status <> 2", "%"+keyword+"%").Find(&users1).RowsAffected
	res := DB.Where("nick_name like ? and status <> 2", "%"+keyword+"%").Limit(10).Offset(int(10 * (pagenum - 1))).Find(&users)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return users, count, nil
}
func GetUserCountByID(userid int64) (int64, int64, int64, int64) {
	usercount := entity.UserCount{}
	DB.Where("user_id = ?", userid).First(&usercount)
	return usercount.FollowCount, usercount.FollowerCount, usercount.WorkCount, usercount.FollowerCount
}

func FollowUser(userid int64, followid int64, status int64) error {
	follow := entity.Follow{FollowerID: uint(userid), FollowID: uint(followid)}
	if status == 0 {
		res := DB.Where("follow_id = ? and follower_id = ?", followid, userid).First(&follow)
		if res.RowsAffected > 0 {
			return errors.New("已关注")
		}
		res = DB.Create(&follow)
		count := entity.UserCount{}
		DB.Where("user_id = ?", userid).First(&count)
		DB.Where("user_id = ?", userid).Update("follow_count", count.FollowCount+1)
		DB.Where("user_id = ?", followid).First(&count)
		DB.Where("user_id = ?", followid).Update("follower_count", count.FollowCount+1)
		if res.Error != nil {
			return res.Error
		}
	} else {
		DB.Where("follow_id = ? and follower_id = ?", followid, userid).First(&follow)
		res := DB.Delete(&follow)
		count := entity.UserCount{}
		DB.Where("user_id = ?", userid).First(&count)
		DB.Where("user_id = ?", userid).Update("follow_count", count.FollowCount-1)
		DB.Where("user_id = ?", followid).First(&count)
		DB.Where("user_id = ?", followid).Update("follower_count", count.FollowCount-1)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func GetFollowUser(userid int64, pagenum int64) ([]entity.User, int64, error) {
	users := []entity.User{}
	users1 := []entity.User{}
	count := DB.Where("id = ?", userid).Find(&users1).RowsAffected
	res := DB.Where("id = ?", userid).Limit(10).Offset(int(10 * (pagenum - 1))).Find(&users)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return users, count, nil
}

func ChangeUserInfo(user entity.User) error {
	res := DB.Model(&entity.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"nick_name": user.NickName,
		"avatar":    user.Avatar,
		"sign":      user.Sign,
		"birth":     user.Birth,
		"sex":       user.Sex,
	})
	if res.Error != nil {
		return res.Error
	}
	if user.Password != "" {
		res = DB.Model(&entity.User{}).Where("id = ?", user.ID).Update("password", user.Password)
	}
	return res.Error
}
