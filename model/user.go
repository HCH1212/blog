package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20)" json:"name"`                      // 用户名
	Age       int    `gorm:"type:int(11)" json:"age"`                           // 年龄
	Email     string `gorm:"type:varchar(20)" json:"email"`                     // 邮箱
	Telephone string `gorm:"type:varchar(20);not null;unique" json:"telephone"` // 手机号，必填且唯一
	Password  string `gorm:"type:varchar(255);not null" json:"password"`        // 密码，必填
	Power     int    `gorm:"type:int(11)" json:"power"`                         // 用户权限 0:普通用户 1:管理员 2:受限用户 默认为0
}

func (User) TableName() string {
	return "user"
}
