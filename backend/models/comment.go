package models

import (
	"errors"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID    uint      `gorm:"index;not null" json:"user_id"`                // 评论者用户id，添加索引且不能为空
	ArticleID uint      `gorm:"index;not null" json:"article_id"`             // 所属文章ID，添加索引且不能为空
	ParentID  *uint     `gorm:"index" json:"parent_id,omitempty"`             // 父评论ID（实现嵌套），添加索引
	Content   string    `gorm:"type:text;not null" json:"content"`            // 评论内容，不能为空
	Replies   []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"` // 子评论
}

func (Comment) TableName() string {
	return "comments"
}

// CreateComment 创建新评论
func CreateComment(db *gorm.DB, comment *Comment) error {
	return db.Create(comment).Error
}

// GetCommentByID 根据ID查询评论
func GetCommentByID(db *gorm.DB, id uint) (*Comment, error) {
	var comment Comment
	err := db.Preload("Replies").First(&comment, id).Error // 预加载子评论
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// DeleteComment 删除评论（增加UserID判断，只有本人才能删除）
func DeleteComment(db *gorm.DB, id uint, userID uint) error {
	// 查询评论是否存在
	var comment Comment
	err := db.First(&comment, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		return err
	}

	// 检查评论的UserID是否与传入的userID一致
	if comment.UserID != userID {
		return errors.New("无权删除该评论")
	}

	// 删除评论
	return db.Delete(&comment).Error
}

// GetCommentsByArticleID 查询文章的所有评论（支持分页）
func GetCommentsByArticleID(db *gorm.DB, articleID uint, page, pageSize int) ([]Comment, error) {
	var comments []Comment
	offset := (page - 1) * pageSize
	err := db.Where("article_id = ? AND parent_id IS NULL", articleID). // 只查询顶级评论
										Offset(offset).Limit(pageSize).
										Preload("Replies"). // 预加载子评论
										Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// GetRepliesByCommentID 查询评论的回复（子评论）
func GetRepliesByCommentID(db *gorm.DB, parentID uint, page, pageSize int) ([]Comment, error) {
	var replies []Comment
	offset := (page - 1) * pageSize
	err := db.Where("parent_id = ?", parentID).
		Offset(offset).Limit(pageSize).
		Find(&replies).Error
	if err != nil {
		return nil, err
	}
	return replies, nil
}

// CheckParentCommentExists 检查父评论是否存在
func CheckParentCommentExists(db *gorm.DB, parentID uint) (bool, error) {
	var comment Comment
	err := db.First(&comment, parentID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 父评论不存在
		}
		return false, err // 其他错误
	}
	return true, nil // 父评论存在
}
