package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Path   string `json:"path"`                 // 图片路径
	Hash   string `json:"hash"`                 // 图片hash
	Name   string `gorm:"size:128" json:"name"` // 图片名称
	Suffix string `gorm:"size:8" json:"suffix"` // 文件后缀
}

func (Image) TableName() string {
	return "image"
}
