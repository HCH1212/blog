package api

import (
	"context"
	"github.com/HCH1212/blog/backend/dao/sqlite"
	"github.com/HCH1212/blog/backend/models"
	"github.com/HCH1212/blog/backend/service"
	"github.com/HCH1212/utils/resp"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// AddComment 添加评论
func AddComment(ctx context.Context, c *app.RequestContext) {
	var req models.Comment
	if err := c.BindAndValidate(&req); err != nil {
		resp.Fail(c, "参数错误", nil)
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		resp.Fail(c, "用户未登录", nil)
		return
	}
	req.UserID = userID.(uint)

	// 调用 Service 层
	if err := service.AddCommentService(sqlite.DB, &req); err != nil {
		resp.FailButServer(c, "添加评论失败", nil)
		return
	}

	resp.Success(c, "添加评论成功", req)
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, c *app.RequestContext) {
	userID, exists := c.Get("user_id")
	if !exists {
		resp.Fail(c, "用户未登录", nil)
		return
	}

	// 获取评论ID
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Fail(c, "评论ID错误", nil)
		return
	}

	// 调用 Service 层
	if err = service.DeleteCommentService(sqlite.DB, uint(id), userID.(uint)); err != nil {
		resp.FailButServer(c, "删除评论失败", nil)
		return
	}

	resp.Success(c, "删除评论成功", nil)
}

// ListComments 查询文章的所有评论（支持分页）
func ListComments(ctx context.Context, c *app.RequestContext) {
	// 获取文章ID
	articleIDStr := c.Query("article_id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		resp.Fail(c, "文章ID错误", nil)
		return
	}

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
	comments, err := service.ListCommentsService(sqlite.DB, uint(articleID), page, pageSize)
	if err != nil {
		resp.FailButServer(c, "查询评论失败", nil)
		return
	}

	resp.Success(c, "查询评论成功", comments)
}

// GetReplies 查询评论的回复（子评论）
func GetReplies(ctx context.Context, c *app.RequestContext) {
	// 获取父评论ID
	parentIDStr := c.Query("parent_id")
	parentID, err := strconv.Atoi(parentIDStr)
	if err != nil {
		resp.Fail(c, "父评论ID错误", nil)
		return
	}

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
	replies, err := service.GetRepliesService(sqlite.DB, uint(parentID), page, pageSize)
	if err != nil {
		resp.FailButServer(c, "查询回复失败", nil)
		return
	}

	resp.Success(c, "查询回复成功", replies)
}
