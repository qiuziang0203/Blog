package mysql

import "Blog/entity"

func GetPending(pagenum int64) ([]entity.Blog, int64, error) {
	blogs := []entity.Blog{}
	blogs1 := []entity.Blog{}
	count := DB.Where("status = 1").Find(&blogs1).RowsAffected
	res := DB.Where("status = 1").Limit(10).Offset(int((pagenum - 1) * 10)).Find(&blogs)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return blogs, count, nil
}
func ApprovalBlog(id int64) error {
	res := DB.Model(&entity.Blog{}).Where("id = ?", id).Update("status", 2)
	blog := entity.Blog{}
	DB.Where("id = ?", id).First(&blog)
	count := entity.UserCount{}
	DB.Where("user_id = ?", blog.UserID).First(&count)
	DB.Update("work_count", count.WorkCount+1)
	return res.Error
}
func ForbiddenUser(id int64, status int64) error {
	res := DB.Model(&entity.User{}).Where("id = ?", id).Update("status", status)
	return res.Error
}

func GetForbidden(pagenum int64) ([]entity.User, int64, error) {
	user := []entity.User{}
	user1 := []entity.User{}
	count := DB.Where("status = 2").Find(&user1).RowsAffected
	res := DB.Where("status = 2").Limit(10).Offset(int((pagenum - 1) * 10)).Find(&user)
	if res.Error != nil {
		return nil, 0, res.Error
	}
	return user, count, nil
}
func AddAnnouncement(announcement entity.Announcement) {
	DB.Create(&announcement)
}
func DelAnnouncement(id int64) {
	DB.Delete(&entity.Announcement{}, id)
}
