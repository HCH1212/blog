package sqlite

import (
	"errors"
	"github.com/HCH1212/blog/backend/models"
	"github.com/HCH1212/utils/password"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	NAME     = "admin"
	PASSWORD = "admin"
)

// Init 初始化 SQLite 数据库
func Init() {
	// 连接到 SQLite 数据库，如果文件不存在，会自动创建名为 test.db 的数据库文件
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect database:", err)
	}

	// 自动迁移模式，根据定义的结构体创建数据库表
	err = db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Comment{},
	)
	if err != nil {
		logrus.Fatal("failed to migrate schema:", err)
	}

	DB = db

	// 初始化管理员用户
	initAdminUser()
}

// initAdminUser 初始化管理员用户
func initAdminUser() {
	// 检查是否已存在管理员用户
	var user models.User
	result := DB.Where("username = ?", "admin").First(&user)

	// 如果不存在管理员用户，则创建
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			name := NAME
			passwordHashed, _ := password.HashPassword(PASSWORD)
			adminUser := models.User{
				Username: name,
				Password: passwordHashed,
				Role:     1,
			}
			// 创建管理员用户
			if err := DB.Create(&adminUser).Error; err != nil {
				logrus.Fatal("failed to create admin user:", err)
			}
		} else {
			logrus.Fatal("db error")
		}
	}
}
