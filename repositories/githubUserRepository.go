package repositories

import (
	"github.com/RocsSun/rocs/model"
	"github.com/jinzhu/gorm"
)

// GithubUserRepository 空
type GithubUserRepository struct{}

// NewGithubUserRepository GithubUserRepository
func NewGithubUserRepository() *GithubUserRepository {
	return &GithubUserRepository{}
}

// Get 根据Id获取
func (r *GithubUserRepository) Get(db *gorm.DB, id int64) *model.GithubUser {
	user := &model.GithubUser{}
	if err := db.First(user, "id = ?", id).Error; err != nil {
		return nil
	}
	return user
}

// Take gorm Take
func (r *GithubUserRepository) Take(db *gorm.DB, where ...interface{}) *model.GithubUser {
	user := &model.GithubUser{}
	if err := db.Take(user, where...).Error; err != nil {
		return nil
	}
	return user
}

// Create 创建GitHub用户
func (r *GithubUserRepository) Create(db *gorm.DB, u *model.GithubUser) (err error) {
	err = db.Create(u).Error
	return
}

// Update 创建GitHub用户
func (r *GithubUserRepository) Update(db *gorm.DB, u *model.GithubUser) (err error) {
	err = db.Save(u).Error
	return
}

// Updates 更新
func (r *GithubUserRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.GithubUser{}).Where("id = ?", id).Updates(columns).Error
	return
}

// UpdateColumn 更新部分
func (r *GithubUserRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.GithubUser{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

// Delete 删除
func (r *GithubUserRepository) Delete(db *gorm.DB, id int64) {
	db.Model(&model.GithubUser{}).Delete("id", id)
}
