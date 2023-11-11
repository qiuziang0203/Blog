package mysql

import (
	"Blog/entity"
	"errors"
)

func CreateBlog(blog entity.Blog, p []string) (entity.Blog, error) {
	tx := DB.Begin()
	res := tx.Create(&blog)
	if res.Error != nil {
		tx.Rollback()
		return blog, res.Error
	}
	for i := 0; i < len(p); i++ {
		e := entity.Type{}
		blogType := entity.BlogType{}
		res := tx.Where("type_name = ?", p[i]).First(&e)
		if res.Error != nil {
			e.TypeName = p[i]
			tx.Create(&e)
			blogType.BlogID = blog.ID
			blogType.TypeID = e.ID
			res = tx.Create(&blogType)
		} else {
			blogType.BlogID = blog.ID
			blogType.TypeID = e.ID
			res = tx.Create(&blogType)
		}
		if res.Error != nil {
			tx.Rollback()
			return blog, res.Error
		}
	}
	tx.Commit()
	return blog, res.Error
}
func GetBlogByType(typename string, way string, pagenum int) ([]entity.Blog, int64, error) {
	typee := entity.Type{}
	res := DB.Where("type_name = ?", typename).First(&typee)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	blogTypes := []entity.BlogType{}
	blogTypes1 := []entity.BlogType{}
	blogs := []entity.Blog{}
	count := DB.Joins("JOIN blogs ON blogs.id = blog_types.blog_id").Where("blog_types.type_id = ? and blogs.status = ?", typee.ID, 2).Find(&blogTypes1).RowsAffected
	res = DB.Joins("JOIN blogs ON blogs.id = blog_types.blog_id").Where("blog_types.type_id = ? and blogs.status = ?", typee.ID, 2).Order(way).Limit(10).Offset(10 * (pagenum - 1)).Find(&blogTypes)
	for i := 0; i < len(blogTypes); i++ {
		blog := entity.Blog{}
		res = DB.Where("id = ? and status = 2", blogTypes[i].BlogID).First(&blog)
		if res.Error == nil {
			blogs = append(blogs, blog)
		}
	}
	return blogs, count, nil
}
func GetBlogByKeyword(keyword string, way string, pagenum int64) ([]entity.Blog, int64, error) {
	blogs := []entity.Blog{}
	blogs1 := []entity.Blog{}
	count := DB.Where("title like ? and status = 2", "%"+keyword+"%").Find(&blogs1).RowsAffected
	res := DB.Where("title like ? and status = 2", "%"+keyword+"%").Order(way).Limit(10).Offset(int(10 * (pagenum - 1))).Find(&blogs)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return blogs, count, nil
}

func GetUserLike(userid int64) ([]int, error) {
	likes := []entity.Like{}
	ids := []int{}
	res := DB.Where("user_id = ?", userid).Find(&likes)
	if res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(likes); i++ {
		ids = append(ids, int(likes[i].BlogID))
	}
	return ids, nil
}

func GetUserFavor(userid int64) ([]int, error) {
	favors := []entity.Favor{}
	ids := []int{}
	res := DB.Where("user_id = ?", userid).Find(&favors)
	if res.Error != nil {
		return nil, res.Error
	}
	for i := 0; i < len(favors); i++ {
		ids = append(ids, int(favors[i].BlogID))
	}
	return ids, nil
}

// 赞，收藏，评论
func GetBlogCount(blogid int64) (int, int, int, error) {
	blog := entity.Blog{}
	res := DB.Where("id = ?", blogid).First(&blog)
	if res.Error != nil {
		return 0, 0, 0, res.Error
	}
	return blog.LikeNum, blog.FavorNum, blog.CommentNum, nil
}

func GetBlogByID(blogid int64) entity.Blog {
	blog := entity.Blog{}
	DB.Where("id = ?", blogid).First(&blog)
	return blog
}
func GetBlogType(blogid int64) []string {
	blogType := []entity.BlogType{}
	var t []string
	DB.Where("blog_id = ?", blogid).Find(&blogType)
	for i := 0; i < len(blogType); i++ {
		typee := entity.Type{}
		DB.Where("id = ?", blogType[i].TypeID).First(&typee)
		t = append(t, typee.TypeName)
	}
	return t
}
func CommentBlog(comment entity.Comment) error {
	res := DB.Create(&comment)
	return res.Error
}
func ReComment(recomment entity.ReComment) error {
	comment := entity.Comment{}
	DB.Where("id = ?", recomment.CommentID).First(&comment)
	DB.Model(&entity.Comment{}).Where("id = ?", recomment.CommentID).Update("re_comment_num", comment.ReCommentNum+1)
	res := DB.Create(&recomment)
	return res.Error
}
func GetBlogComment(blogid int64) []entity.Comment {
	comments := []entity.Comment{}
	DB.Where("blog_id = ?", blogid).Find(&comments)
	return comments
}
func GetReComment(commentid int64) []entity.ReComment {
	recomments := []entity.ReComment{}
	DB.Where("comment_id = ?", commentid).Find(&recomments)
	return recomments
}
func LikeComment(commentid int64, userid int64) error {
	like := entity.CommentLike{}
	res := DB.Where("comment_id = ? and userid = ?", commentid, userid).First(&like)
	if res.RowsAffected > 0 {
		return errors.New("已点赞")
	}
	comment := entity.Comment{}
	DB.Where("id = ?", commentid).First(&comment)
	DB.Model(&entity.Comment{}).Where("id = ?", commentid).Update("like_num", comment.LikeNum+1)
	like.CommentID = uint(commentid)
	like.UserID = uint(userid)
	DB.Create(&like)
	return nil
}

func DelLikeComment(commentid int64, userid int64) error {
	like := entity.CommentLike{}
	res := DB.Where("comment_id = ? and userid = ?", commentid, userid).First(&like)
	if res.RowsAffected == 0 {
		return errors.New("未点赞")
	}
	comment := entity.Comment{}
	DB.Where("id = ?", commentid).First(&comment)
	DB.Model(&entity.Comment{}).Where("id = ?", commentid).Update("like_num", comment.LikeNum-1)
	DB.Delete(&entity.CommentLike{}, like.ID)
	return nil
}
func GetDraft(userid int64, pagenum int64) ([]entity.Blog, int64, error) {
	blogs := []entity.Blog{}
	blogs1 := []entity.Blog{}
	count := DB.Where("user_id = ? and status = 0", userid).Find(&blogs1).RowsAffected
	res := DB.Where("user_id = ? and status = 0", userid).Limit(10).Offset(int(10 * (pagenum - 1))).Find(&blogs)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return blogs, count, nil
}
func GetUserBlog(userid int64, pagenum int64) ([]entity.Blog, int64, error) {
	blogs := []entity.Blog{}
	blogs1 := []entity.Blog{}
	count := DB.Where("user_id = ? and status = 2", userid).Find(&blogs1).RowsAffected
	res := DB.Where("user_id = ? and status = 2", userid).Limit(10).Offset(int(10 * (pagenum - 1))).Find(&blogs)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return blogs, count, nil
}

func DelType(blogid int64) {
	blogType := []entity.BlogType{}
	DB.Where("blog_id = ?", blogid).Find(&blogType)
	for i := 0; i < len(blogType); i++ {
		DB.Delete(&entity.BlogType{}, blogType[i].ID)
	}
}
func AddType(blogid int64, p []string) {
	tx := DB.Begin()
	for i := 0; i < len(p); i++ {
		e := entity.Type{}
		blogType := entity.BlogType{}
		res := tx.Where("type_name = ?", p[i]).First(&e)
		if res.Error != nil {
			e.TypeName = p[i]
			tx.Create(&e)
			blogType.BlogID = uint(blogid)
			blogType.TypeID = e.ID
			res = tx.Create(&blogType)
		} else {
			blogType.BlogID = uint(blogid)
			blogType.TypeID = e.ID
			res = tx.Create(&blogType)
		}
		if res.Error != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
}

func ChangeBlogInfo(blog entity.Blog) error {
	res := DB.Model(&entity.Blog{}).Where("id = ?", blog.ID).Updates(map[string]interface{}{
		"title":  blog.Title,
		"text":   blog.Text,
		"status": blog.Status,
	})
	return res.Error
}

func GetHotBlog() ([]entity.Blog, int, error) {
	blogs := []entity.Blog{}
	blogs1 := []entity.Blog{}
	count := DB.Where("status = ?", 2).Find(&blogs1).RowsAffected
	res := DB.Where("status = ?", 2).Order("hot desc").Limit(10).Offset(0).Find(&blogs)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return blogs, int(count), nil
}

func GetAllType() []entity.Type {
	t := []entity.Type{}
	DB.Find(&t)
	return t
}
func GetAnnouncement() []entity.Announcement {
	announcement := []entity.Announcement{}
	DB.Find(&announcement)
	return announcement
}
func DelBlog(blogid int64) error {
	res := DB.Delete(&entity.Blog{}, blogid)
	return res.Error
}
