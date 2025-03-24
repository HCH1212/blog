package service

import (
	"errors"
	"github.com/HCH1212/blog/backend/models"
	"gorm.io/gorm"
)

// AddArticleService 添加文章
func AddArticleService(db *gorm.DB, article *models.Article) error {
	return models.CreateArticle(db, article)
}

// UpdateArticleService 更新文章
func UpdateArticleService(db *gorm.DB, id uint, updatedArticle *models.Article) error {
	return models.UpdateArticle(db, id, updatedArticle)
}

// DeleteArticleService 删除文章
func DeleteArticleService(db *gorm.DB, id uint) error {
	return models.DeleteArticle(db, id)
}

// ListArticlesService 查询文章列表（支持分页）
func ListArticlesService(db *gorm.DB, page, pageSize int) ([]models.Article, int64, error) {
	return models.GetAllArticles(db, page, pageSize)
}

// GetArticleService 根据ID查询文章
func GetArticleService(db *gorm.DB, id uint) (*models.Article, error) {
	return models.GetArticleByID(db, id)
}

// AddFavoriteService 添加收藏
func AddFavoriteService(db *gorm.DB, userID uint, articleID uint) error {
	return models.AddFavorite(db, userID, articleID)
}

// RemoveFavoriteService 取消收藏
func RemoveFavoriteService(db *gorm.DB, userID uint, articleID uint) error {
	return models.RemoveFavorite(db, userID, articleID)
}

// SearchService 搜索
func SearchService(db *gorm.DB, q string) ([]models.Article, error) {
	res, err := models.SearchArticles(db, q)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return res, err
}

func GetArticlesByTagService(db *gorm.DB, tag string) ([]models.Article, error) {
	return models.GetArticlesByTag(db, tag)
}

func GetTagsService(db *gorm.DB) ([]string, error) {
	return models.GetAllTags(db)
}
