package timedtask

import (
	"Blog/dao/mysql"
	"Blog/dao/redis"
	"Blog/entity"
	"context"
	"encoding/json"
	"strconv"
)

func LikeBlog() error {
	ctx := context.Background()
	var blogids []int
	vals, err := redis.Rdb.HGet(ctx, "blog_count_change", "like").Result()
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(vals), &blogids)
	for i := 0; i < len(blogids); i++ {
		likestr, _ := redis.Rdb.HGet(ctx, "blog_count::blog_id"+strconv.FormatInt(int64(blogids[i]), 10), "like").Result()
		like, _ := strconv.Atoi(likestr)
		res := mysql.DB.Model(&entity.Blog{}).Where("id = ?", blogids[i]).Update("like_num", like)
		if res.Error != nil {
			return err
		}
		likes, favors, _, _ := redis.GetBlogCount(int64(blogids[i]))
		mysql.DB.Model(&entity.Blog{}).Where("id = ?", blogids[i]).Update("hot", likes*2+favors*4)
	}
	redis.Rdb.HDel(ctx, "blog_count_change", "like")
	return nil
}
func FavorBlog() error {
	ctx := context.Background()
	var blogids []int
	vals, err := redis.Rdb.HGet(ctx, "blog_count_change", "favor").Result()
	if err != nil {
		return err
	}
	json.Unmarshal([]byte(vals), &blogids)
	for i := 0; i < len(blogids); i++ {
		likestr, _ := redis.Rdb.HGet(ctx, "blog_count::blog_id"+strconv.FormatInt(int64(blogids[i]), 10), "favor").Result()
		like, _ := strconv.Atoi(likestr)
		res := mysql.DB.Model(&entity.Blog{}).Where("id = ?", blogids[i]).Update("favor_num", like)
		if res.Error != nil {
			return err
		}
		likes, favors, _, _ := redis.GetBlogCount(int64(blogids[i]))
		mysql.DB.Model(&entity.Blog{}).Where("id = ?", blogids[i]).Update("hot", likes*2+favors*4)
	}
	redis.Rdb.HDel(ctx, "blog_count_change", "favor")
	return nil
}
