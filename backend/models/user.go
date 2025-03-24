package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string    `gorm:"unique;not null" json:"username"`                      // 用户名，唯一且不能为空
	Password  string    `gorm:"not null" json:"password"`                             // 用户密码，不能为空
	Role      uint8     `gorm:"default:0" json:"role,omitempty"`                      // 用户角色：admin:1 或 user:0
	Favorites []Article `gorm:"many2many:user_favorites;" json:"favorites,omitempty"` // 用户收藏的文章，多对多关系
}

func (User) TableName() string {
	return "users"
}

// GetUserByName 通过用户名查询用户
func GetUserByName(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// GetUserByID 通过id查询用户
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUserFavoritesByID 根据用户ID查询收藏的文章
func GetUserFavoritesByID(db *gorm.DB, userID uint) (*User, error) {
	var user User
	// 使用 Preload 加载 Favorites 关联
	err := db.Preload("Favorites").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserFavoritesByIDWithPagination 根据用户ID分页查询收藏的文章
func GetUserFavoritesByIDWithPagination(db *gorm.DB, userID uint, page, pageSize int) (*User, error) {
	var user User
	offset := (page - 1) * pageSize
	// 使用 Preload 加载 Favorites 关联，并分页
	err := db.Preload("Favorites", func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize) // 分页
	}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
