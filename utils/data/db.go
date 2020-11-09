package data

import "github.com/jinzhu/gorm"

var db *gorm.DB

// DB 获取数据库连接
func DB() *gorm.DB {
	return db
}
