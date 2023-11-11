package mysql

import (
	"Blog/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(mysql.Open(utils.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败，error=" + err.Error())
	}
	DB = db
}
