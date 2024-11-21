// 系统管理
package api

import (
	"blog/resp"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	resp.Success(c, "1111111111", nil)
}
