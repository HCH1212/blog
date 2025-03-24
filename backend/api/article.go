package api

import (
	"context"
	"github.com/HCH1212/blog/backend/dao/sqlite"
	"github.com/HCH1212/blog/backend/models"
	"github.com/HCH1212/blog/backend/service"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(ctx context.Context, c *app.RequestContext) {
	var req models.Article
	if err := c.BindAndValidate(&req); err != nil {
		resp.Fail(c, "参数错误", nil)
		return
	}

	// 调用 Service 层
	if err := service.AddArticleService(sqlite.DB, &req); err != nil {
		resp.FailButServer(c, "添加文章失败", nil)
		return
	}

	resp.Success(c, "添加文章成功", req)
}

// UpdateArticle 更新文章
func UpdateArticle(ctx context.Context, c *app.RequestContext) {
	var req models.Article
	if err := c.BindAndValidate(&req); err != nil {
		resp.Fail(c, "参数错误", nil)
		return
	}

	// 获取文章ID
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Fail(c, "文章ID错误", nil)
		return
	}

	// 调用 Service 层
	if err = service.UpdateArticleService(sqlite.DB, uint(id), &req); err != nil {
		resp.FailButServer(c, "更新文章失败", nil)
		return
	}

	resp.Success(c, "更新文章成功", req)
}

// DeleteArticle 删除文章
func DeleteArticle(ctx context.Context, c *app.RequestContext) {
	// 获取文章ID
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Fail(c, "文章ID错误", nil)
		return
	}

	// 调用 Service 层
	if err = service.DeleteArticleService(sqlite.DB, uint(id)); err != nil {
		resp.FailButServer(c, "删除文章失败", nil)
		return
	}

	resp.Success(c, "删除文章成功", nil)
}

// ListArticles 查询文章列表（支持分页）
func ListArticles(ctx context.Context, c *app.RequestContext) {
	// 获取分页参数
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // 默认第一页
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10 // 默认每页10条
	}

	// 调用 Service 层
	articles, total, err := service.ListArticlesService(sqlite.DB, page, pageSize)
	if err != nil {
		resp.FailButServer(c, "查询文章列表失败", nil)
		return
	}

	resp.Success(c, "查询文章列表成功", utils.H{"articles": articles, "total": total})
}

// GetArticleByID 根据ID查询文章
func GetArticle(ctx context.Context, c *app.RequestContext) {
	// 获取文章ID
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Fail(c, "文章ID错误", nil)
		return
	}

	// 调用 Service 层
	article, err := service.GetArticleService(sqlite.DB, uint(id))
	if err != nil {
		resp.FailButServer(c, "查询文章失败", nil)
		return
	}

	resp.Success(c, "查询文章成功", article)
}

// AddFavorite 添加收藏
func AddFavorite(ctx context.Context, c *app.RequestContext) {
	// 获取用户ID（从上下文中获取，假设已经通过认证）
	userID, exists := c.Get("user_id")
	if !exists {
		resp.Fail(c, "用户未登录", nil)
		return
	}

	// 获取文章ID
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		resp.Fail(c, "文章ID错误", nil)
		return
	}

	// 调用 Service 层
	if err = service.AddFavoriteService(sqlite.DB, userID.(uint), uint(articleID)); err != nil {
		resp.FailButServer(c, "添加收藏失败", nil)
		return
	}

	resp.Success(c, "添加收藏成功", nil)
}

// RemoveFavorite 取消收藏
func RemoveFavorite(ctx context.Context, c *app.RequestContext) {
	// 获取用户ID（从上下文中获取，假设已经通过认证）
	userID, exists := c.Get("user_id")
	if !exists {
		resp.Fail(c, "用户未登录", nil)
		return
	}

	// 获取文章ID
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		resp.Fail(c, "文章ID错误", nil)
		return
	}

	// 调用 Service 层
	if err = service.RemoveFavoriteService(sqlite.DB, userID.(uint), uint(articleID)); err != nil {
		resp.FailButServer(c, "取消收藏失败", nil)
		return
	}

	resp.Success(c, "取消收藏成功", nil)
}

func Search(ctx context.Context, c *app.RequestContext) {
	q := c.Query("q")

	// 调用 Service 层
	articles, err := service.SearchService(sqlite.DB, q)
	if err != nil {
		resp.FailButServer(c, "搜索文章失败", nil)
		return
	}

	resp.Success(c, "成功", articles)
}

// GetArticlesByTagHandler 通过标签查询文章的处理函数
func GetArticlesByTag(ctx context.Context, c *app.RequestContext) {
	tag := c.Query("tag")
	if tag == "" {
		resp.Fail(c, "标签参数缺失", nil)
		return
	}

	articles, err := service.GetArticlesByTagService(sqlite.DB, tag)
	if err != nil {
		resp.FailButServer(c, "查询文章失败", nil)
		return
	}

	resp.Success(c, "查询文章成功", articles)
}

func GetTags(ctx context.Context, c *app.RequestContext) {
	tags, err := service.GetTagsService(sqlite.DB)
	if err != nil {
		resp.FailButServer(c, "获取所有标签失败", nil)
		return
	}

	resp.Success(c, "获取所有标签成功", tags)
}
