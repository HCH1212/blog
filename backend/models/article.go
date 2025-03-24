package models

import (
	"gorm.io/gorm"
	"strings"
)

type Article struct {
	gorm.Model
	Title      string    `gorm:"type:varchar(255);unique;not null" json:"title"` // 文章标题，长度不超过255，唯一且不能为空
	Content    string    `gorm:"type:text;not null" json:"content"`              // 文章内容，不能为空
	Tags       string    `gorm:"type:varchar(255)" json:"tags,omitempty"`        // 逗号分隔的标签
	CoverImage string    `gorm:"type:varchar(255)" json:"cover_image,omitempty"` // 封面图URL
	Comments   []Comment `gorm:"foreignKey:ArticleID" json:"comments,omitempty"` // 文章的评论
}

func (Article) TableName() string {
	return "articles"
}

// CreateArticle 创建新文章
func CreateArticle(db *gorm.DB, article *Article) error {
	return db.Create(article).Error
}

// GetArticleByID 根据ID查询文章
func GetArticleByID(db *gorm.DB, id uint) (*Article, error) {
	var article Article
	err := db.Preload("Comments").First(&article, id).Error // 预加载评论
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// UpdateArticle 更新文章
func UpdateArticle(db *gorm.DB, id uint, updatedArticle *Article) error {
	var article Article
	// 先查询文章是否存在
	err := db.First(&article, id).Error
	if err != nil {
		return err
	}
	// 更新字段
	return db.Model(&article).Updates(updatedArticle).Error
}

// DeleteArticle 删除文章
func DeleteArticle(db *gorm.DB, id uint) error {
	return db.Delete(&Article{}, id).Error
}

// GetAllArticles 查询所有文章（支持分页）
func GetAllArticles(db *gorm.DB, page, pageSize int) ([]Article, int64, error) {
	var articles []Article
	var total int64

	// 先查询总记录数
	err := db.Model(&Article{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = db.Offset(offset).Limit(pageSize).Find(&articles).Error
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

// AddFavorite 添加收藏
func AddFavorite(db *gorm.DB, userID uint, articleID uint) error {
	// 查询用户
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	// 查询文章
	var article Article
	if err := db.First(&article, articleID).Error; err != nil {
		return err
	}

	// 添加收藏
	return db.Model(&user).Association("Favorites").Append(&article)
}

// RemoveFavorite 取消收藏
func RemoveFavorite(db *gorm.DB, userID uint, articleID uint) error {
	// 查询用户
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	// 查询文章
	var article Article
	if err := db.First(&article, articleID).Error; err != nil {
		return err
	}

	// 取消收藏
	return db.Model(&user).Association("Favorites").Delete(&article)
}

// SearchArticles 根据搜索关键词模糊查询文章
func SearchArticles(db *gorm.DB, q string) ([]Article, error) {
	var articles []Article

	if q == "" {
		return nil, nil // 如果没有提供搜索关键词，返回空结果
	}

	// 构造模糊搜索的查询条件
	searchPattern := "%" + q + "%"
	err := db.Where("title LIKE ? OR content LIKE ?", searchPattern, searchPattern).Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticlesByTag 根据标签查询文章
func GetArticlesByTag(db *gorm.DB, tag string) ([]Article, error) {
	var articles []Article
	err := db.Where("tags LIKE ?", "%"+tag+"%").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// GetAllTags 获取所有标签
func GetAllTags(db *gorm.DB) ([]string, error) {
	var articles []Article
	if err := db.Select("tags").Find(&articles).Error; err != nil {
		return nil, err
	}

	tagMap := make(map[string]struct{})
	for _, article := range articles {
		if article.Tags != "" {
			tags := strings.Split(article.Tags, ",")
			for _, tag := range tags {
				tag = strings.TrimSpace(tag)
				if tag != "" {
					tagMap[tag] = struct{}{}
				}
			}
		}
	}

	var allTags []string
	for tag := range tagMap {
		allTags = append(allTags, tag)
	}

	return allTags, nil
}
