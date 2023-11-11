package redis

import (
	"Blog/dao/daoutils"
	"Blog/dao/mysql"
	"Blog/entity"
	"context"
	"encoding/json"
	"errors"
	"slices"
	"strconv"
	"time"
)

// 将博客存入redis
func CreateBlog(blog entity.Blog, t []string) error {
	ctx := context.Background()
	m := make(map[string]interface{})

	m["id"] = blog.ID
	m["user_id"] = blog.UserID
	m["title"] = blog.Title
	m["text"] = blog.Text
	m["status"] = blog.Status
	bs, _ := json.Marshal(t)
	m["type"] = bs
	m["time"] = blog.UpdatedAt
	err := Rdb.HMSet(ctx, "blog_info::blog_id"+strconv.FormatInt(int64(blog.ID), 10), m).Err()
	Rdb.Expire(ctx, "blog_info::blog_id"+strconv.FormatInt(int64(blog.ID), 10), time.Hour*24*7)
	return err
}

func GetUserLike(userid int64) ([]int, error) {
	ctx := context.Background()
	var ids []int
	vals, err := Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), "likes").Result()

	if err != nil || vals == "" {
		ids, err = mysql.GetUserLike(userid)
		if err != nil {
			return nil, err
		}
		bs, _ := json.Marshal(ids)
		m := make(map[string]interface{})
		m["likes"] = bs
		err = Rdb.HMSet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), m).Err()
		Rdb.Expire(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), time.Hour*24*7)
	} else {
		err = json.Unmarshal([]byte(vals), &ids)
	}
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func GetUserFavor(userid int64) ([]int, error) {
	ctx := context.Background()
	var ids []int
	vals, err := Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), "favors").Result()
	if err != nil {
		ids, err = mysql.GetUserFavor(userid)
		if err != nil {
			return nil, err
		}
		bs, _ := json.Marshal(ids)
		m := make(map[string]interface{})
		m["favors"] = bs
		err = Rdb.HMSet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), m).Err()
		Rdb.Expire(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), time.Hour*24*7)
	} else {
		err = json.Unmarshal([]byte(vals), &ids)
	}
	if err != nil {
		return nil, err
	}
	return ids, nil
}
func LikeBlog(userid int64, blogid int64, status int64) error {
	ctx := context.Background()
	if status == 0 {
		flag := daoutils.IsLike(userid, blogid)
		if flag == true {
			return errors.New("已点赞")
		}
		var addids []int
		vals, _ := Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), "adds").Result()
		json.Unmarshal([]byte(vals), &addids)
		if slices.Contains(addids, int(blogid)) == false {
			addids = append(addids, int(blogid))
		}
		bs, _ := json.Marshal(addids)
		m := make(map[string]interface{})
		m["adds"] = bs
		Rdb.HMSet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), m)
		Rdb.Expire(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), time.Hour*24*7)
		Rdb.HIncrBy(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "like", 1)
	} else if status == 1 {
		flag := daoutils.IsLike(userid, blogid)
		if flag == false {
			return errors.New("未点赞")
		}
		var addids []int
		vals, _ := Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), "adds").Result()
		json.Unmarshal([]byte(vals), &addids)
		for i := 0; i < len(addids); i++ {
			if addids[i] == int(blogid) {
				addids[i] = 0
				return nil
			}
		}
		var delids []int
		vals, _ = Rdb.HGet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), "dels").Result()
		json.Unmarshal([]byte(vals), &delids)
		if slices.Contains(delids, int(blogid)) == false {
			delids = append(delids, int(blogid))
		}
		bs, _ := json.Marshal(delids)
		m := make(map[string]interface{})
		m["dels"] = bs
		Rdb.HMSet(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), m)
		Rdb.Expire(ctx, "user_like::user_id"+strconv.FormatInt(userid, 10), time.Hour*24*7)
		Rdb.HIncrBy(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "like", -1)
	}
	//记录点赞的用户id
	var like []int
	vals, _ := Rdb.HGet(ctx, "user_count_change", "like").Result()
	json.Unmarshal([]byte(vals), &like)
	if slices.Contains(like, int(userid)) == false {
		like = append(like, int(userid))
	}
	bs, _ := json.Marshal(like)
	m := make(map[string]interface{})
	m["like"] = bs
	Rdb.HMSet(ctx, "user_count_change", m)
	//记录变化的博客id
	var likeCount []int
	vals, _ = Rdb.HGet(ctx, "blog_count_change", "like").Result()
	json.Unmarshal([]byte(vals), &likeCount)
	if slices.Contains(likeCount, int(blogid)) == false {
		likeCount = append(likeCount, int(blogid))
	}
	bs, _ = json.Marshal(likeCount)
	m = make(map[string]interface{})
	m["like"] = bs
	Rdb.HMSet(ctx, "blog_count_change", m)
	return nil
}
func FavorBlog(userid int64, blogid int64, status int64) error {
	ctx := context.Background()
	if status == 0 {
		flag := daoutils.IsFavor(userid, blogid)
		if flag == true {
			return errors.New("已收藏")
		}
		var addids []int
		vals, _ := Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), "adds").Result()
		json.Unmarshal([]byte(vals), &addids)
		if slices.Contains(addids, int(blogid)) == false {
			addids = append(addids, int(blogid))
		}
		bs, _ := json.Marshal(addids)
		m := make(map[string]interface{})
		m["adds"] = bs
		Rdb.HMSet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), m)
		Rdb.Expire(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), time.Hour*24*7)
		Rdb.HIncrBy(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "favor", 1)
	} else if status == 1 {
		flag := daoutils.IsFavor(userid, blogid)
		if flag == false {
			return errors.New("未收藏")
		}
		var addids []int
		vals, _ := Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), "adds").Result()
		json.Unmarshal([]byte(vals), &addids)
		for i := 0; i < len(addids); i++ {
			if addids[i] == int(blogid) {
				addids[i] = 0
				return nil
			}
		}
		var delids []int
		vals, _ = Rdb.HGet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), "dels").Result()
		json.Unmarshal([]byte(vals), &delids)
		if slices.Contains(delids, int(blogid)) == false {
			delids = append(delids, int(blogid))
		}
		bs, _ := json.Marshal(delids)
		m := make(map[string]interface{})
		m["dels"] = bs
		Rdb.HMSet(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), m)
		Rdb.Expire(ctx, "user_favor::user_id"+strconv.FormatInt(userid, 10), time.Hour*24*7)
		Rdb.HIncrBy(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "favor", -1)
	}
	//记录收藏的用户id
	var like []int
	vals, _ := Rdb.HGet(ctx, "user_count_change", "favor").Result()
	json.Unmarshal([]byte(vals), &like)
	if slices.Contains(like, int(userid)) == false {
		like = append(like, int(userid))
	}
	bs, _ := json.Marshal(like)
	m := make(map[string]interface{})
	m["favor"] = bs
	Rdb.HMSet(ctx, "user_count_change", m)
	//记录变化的博客id
	var likeCount []int
	vals, _ = Rdb.HGet(ctx, "blog_count_change", "favor").Result()
	json.Unmarshal([]byte(vals), &likeCount)
	if slices.Contains(likeCount, int(blogid)) == false {
		likeCount = append(likeCount, int(blogid))
	}
	bs, _ = json.Marshal(likeCount)
	m = make(map[string]interface{})
	m["favor"] = bs
	Rdb.HMSet(ctx, "blog_count_change", m)
	return nil
}

func GetBlogCount(blogid int64) (int, int, int, error) {
	ctx := context.Background()
	var like int
	var favor int
	var comment int
	vals1, err := Rdb.HGet(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "like").Result()
	vals2, err := Rdb.HGet(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "favor").Result()
	vals3, err := Rdb.HGet(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), "comment").Result()
	if err != nil {
		like, favor, comment, err = mysql.GetBlogCount(blogid)
		if err != nil {
			return 0, 0, 0, err
		}
		m := make(map[string]interface{})
		m["like"] = like
		m["favor"] = favor
		m["comment"] = comment
		err = Rdb.HMSet(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), m).Err()
		Rdb.Expire(ctx, "blog_count::blog_id"+strconv.FormatInt(blogid, 10), time.Hour*24*7)
	}
	like, _ = strconv.Atoi(vals1)
	favor, _ = strconv.Atoi(vals2)
	comment, _ = strconv.Atoi(vals3)
	return like, favor, comment, nil
}

func GetBlogByID(blogid int64) (entity.Blog, []string, error) {
	ctx := context.Background()
	err := Rdb.HExists(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10), "id").Err()
	if err != nil {
		blog := mysql.GetBlogByID(blogid)
		t := mysql.GetBlogType(blogid)
		m := make(map[string]interface{})
		m["id"] = blog.ID
		m["user_id"] = blog.UserID
		m["title"] = blog.Title
		m["text"] = blog.Text
		m["status"] = blog.Status
		bs, _ := json.Marshal(t)
		m["type"] = bs
		m["time"] = blog.UpdatedAt
		Rdb.HMSet(ctx, "blog_info::blog_id"+strconv.FormatInt(int64(blog.ID), 10), m)
		Rdb.Expire(ctx, "blog_info::blog_id"+strconv.FormatInt(int64(blog.ID), 10), time.Hour*24*7)
		return blog, t, nil
	}
	blog := entity.Blog{}
	idstr, err := Rdb.HGet(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10), "id").Result()
	id, _ := strconv.Atoi(idstr)
	blog.ID = uint(id)
	useridstr, err := Rdb.HGet(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10), "user_id").Result()
	userid, _ := strconv.Atoi(useridstr)
	blog.UserID = uint(userid)
	blog.Title, _ = Rdb.HGet(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10), "title").Result()
	blog.Text, _ = Rdb.HGet(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10), "text").Result()
	blog.Status = 2
	blog.LikeNum, blog.FavorNum, blog.CommentNum, _ = GetBlogCount(blogid)
	var types []string
	vals, _ := Rdb.HGet(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10), "type").Result()
	json.Unmarshal([]byte(vals), &types)
	return blog, types, nil
}
func DelBlog(blogid int64) {
	ctx := context.Background()
	Rdb.Del(ctx, "blog_info::blog_id"+strconv.FormatInt(blogid, 10))
}
