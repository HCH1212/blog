package api

import (
	"blog/resp"
	"blog/service"
	"github.com/gin-gonic/gin"
)

func ImageUpdate(c *gin.Context) {
	res, err := service.ImageUpdateService(c)
	if err != nil {
		resp.Fail(c, err.Error(), nil)
		return
	}
	resp.Success(c, "ok", res)
}
