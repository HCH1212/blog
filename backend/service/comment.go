package service

import (
	"errors"
	"github.com/HCH1212/blog/backend/models"
	"gorm.io/gorm"
)

// AddCommentService 添加评论
func AddCommentService(db *gorm.DB, comment *models.Comment) error {
	if comment.ParentID != nil {
		ok, err := models.CheckParentCommentExists(db, *comment.ParentID)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("父评论不存在")
		}
	}
	return models.CreateComment(db, comment)
}

// DeleteCommentService 删除评论
func DeleteCommentService(db *gorm.DB, id, userID uint) error {
	return models.DeleteComment(db, id, userID)
}

// ListCommentsService 查询文章的所有评论（支持分页）
func ListCommentsService(db *gorm.DB, articleID uint, page, pageSize int) ([]models.Comment, error) {
	return models.GetCommentsByArticleID(db, articleID, page, pageSize)
}

// GetRepliesService 查询评论的回复（子评论）
func GetRepliesService(db *gorm.DB, parentID uint, page, pageSize int) ([]models.Comment, error) {
	return models.GetRepliesByCommentID(db, parentID, page, pageSize)
}
