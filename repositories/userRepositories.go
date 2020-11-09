package repositories

import (
	"rocs/model"

	"github.com/jinzhu/gorm"
)

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type UserRepository struct {
}

// NewUserRepository 构建UserRepository指针
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Get 通过ID查询用户
func (r *UserRepository) Get(db *gorm.DB, id int64) *model.User {
	user := &model.User{}
	if err := db.First(user, "id = ?", id).Error; err != nil {
		return nil
	}
	return user
}

// Take gorm.take
func (r *UserRepository) Take(db *gorm.DB, where ...interface{}) *model.User {
	user := &model.User{}
	if err := db.Take(user, where...).Error; err != nil {
		return nil
	}
	return user
}

//func (this *UserRepository) QueryCnd(db *gorm.DB, cnd *simple.QueryCnd) (list []model.User, err error) {
//	err = cnd.DoQuery(db).Find(&list).Error
//	return
//}
//
//func (this *UserRepository) Query(db *gorm.DB, queries *simple.ParamQueries) (list []model.User, paging *simple.Paging) {
//	queries.StartQuery(db).Find(&list)
//	queries.StartCount(db).Model(&model.User{}).Count(&queries.Paging.Total)
//	paging = queries.Paging
//	return
//}

// Create 创建用户
func (r *UserRepository) Create(db *gorm.DB, u *model.User) (err error) {
	err = db.Create(u).Error
	return
}

// Update 更新用户
func (r *UserRepository) Update(db *gorm.DB, u *model.User) (err error) {
	err = db.Save(u).Error
	return
}

// Updates 根据ID更新整行
func (r *UserRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).Update(columns).Error
	return
}

// UpdateColumn 根据ID 更新行中的name对应的value
func (r *UserRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

// Delete 根据ID删除
func (r *UserRepository) Delete(db *gorm.DB, id int64) {
	db.Model(&model.User{}).Delete("id", id)
}

// GetByEmail 通过邮件地址获取
func (r *UserRepository) GetByEmail(db *gorm.DB, email string) *model.User {
	return r.Take(db, "email = ?", email)
}

// GetByUsername 通过用户名地址获取
func (r *UserRepository) GetByUsername(db *gorm.DB, username string) *model.User {
	return r.Take(db, "username = ?", username)
}
