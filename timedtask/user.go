package timedtask

import (
	"Blog/dao/mysql"
	"Blog/dao/redis"
	"Blog/entity"
	"context"
	"encoding/json"
	"slices"
	"strconv"
)

func FavorUser() error {
	ctx := context.Background()
	var userids []int
	vals, err := redis.Rdb.HGet(ctx, "user_count_change", "favor").Result()
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(vals), &userids)
	for i := 0; i < len(userids); i++ {
		var delids, addids, likeids []int
		vals, _ = redis.Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), "favors").Result()
		json.Unmarshal([]byte(vals), &likeids)
		vals, _ = redis.Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), "dels").Result()
		json.Unmarshal([]byte(vals), &delids)
		for j := 0; j < len(delids); j++ {
			likeids = slices.DeleteFunc(likeids, func(k int) bool {
				return k == delids[j]
			})
			like := entity.Favor{}
			mysql.DB.Where("user_id = ? and blog_id = ?", userids[i], delids[j]).First(&like)
			err = mysql.DB.Delete(&like).Error
			if err != nil {
				return err
			}
		}
		vals, _ = redis.Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), "adds").Result()
		json.Unmarshal([]byte(vals), &addids)
		for j := 0; j < len(addids); j++ {
			if addids[j] == 0 {
				continue
			}
			like := entity.Favor{UserID: uint(userids[i]), BlogID: uint(addids[j])}
			like1 := entity.Favor{}
			res := mysql.DB.Where("user_id = ? and blog_id = ?", userids[i], addids[j]).First(&like1)
			if res.RowsAffected > 0 {
				redis.Rdb.HDel(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), "adds")
				continue
			}
			err = mysql.DB.Create(&like).Error
			if err != nil {
				return err
			}
			if slices.Contains(likeids, addids[j]) == false {
				likeids = append(likeids, addids[j])
			}
		}
		bs, _ := json.Marshal(likeids)
		m := make(map[string]interface{})
		m["favors"] = bs
		redis.Rdb.HMSet(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), m)
		redis.Rdb.HDel(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), "adds")
		redis.Rdb.HDel(ctx, "user_favor::user_id"+strconv.FormatInt(int64(userids[i]), 10), "dels")
		favor := []entity.Favor{}
		mysql.DB.Where("user_id = ?", userids[i]).Find(&favor)
		mysql.DB.Model(&entity.UserCount{}).Where("user_id = ?", userids[i]).Update("favor_count", len(favor))
	}
	redis.Rdb.HDel(ctx, "user_count_change", "favor")
	return nil
}
func LikeUser() error {
	ctx := context.Background()
	var userids []int
	vals, err := redis.Rdb.HGet(ctx, "user_count_change", "like").Result()
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(vals), &userids)
	for i := 0; i < len(userids); i++ {
		var delids, addids, likeids []int
		vals, _ = redis.Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), "likes").Result()
		json.Unmarshal([]byte(vals), &likeids)
		vals, _ = redis.Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), "dels").Result()
		json.Unmarshal([]byte(vals), &delids)
		for j := 0; j < len(delids); j++ {
			likeids = slices.DeleteFunc(likeids, func(k int) bool {
				return k == delids[j]
			})
			like := entity.Like{}
			mysql.DB.Where("user_id = ? and blog_id = ?", userids[i], delids[j]).First(&like)
			err = mysql.DB.Delete(&like).Error
			if err != nil {
				return err
			}
		}
		vals, _ = redis.Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), "adds").Result()
		json.Unmarshal([]byte(vals), &addids)
		for j := 0; j < len(addids); j++ {
			if addids[j] == 0 {
				continue
			}
			like := entity.Like{UserID: uint(userids[i]), BlogID: uint(addids[j])}
			like1 := []entity.Like{}
			res := mysql.DB.Where("user_id = ? and blog_id = ?", userids[i], addids[j]).Find(&like1)
			if res.RowsAffected > 0 {
				redis.Rdb.HDel(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), "adds")
				continue
			}
			err = mysql.DB.Create(&like).Error
			if err != nil {
				return err
			}
			if slices.Contains(likeids, addids[j]) == false {
				likeids = append(likeids, addids[j])
			}
		}
		bs, _ := json.Marshal(likeids)
		m := make(map[string]interface{})
		m["likes"] = bs
		redis.Rdb.HSet(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), m)
		redis.Rdb.HDel(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), "adds")
		redis.Rdb.HDel(ctx, "user_like::user_id"+strconv.FormatInt(int64(userids[i]), 10), "dels")
	}
	redis.Rdb.HDel(ctx, "user_count_change", "like")
	return nil
}
