package resp

import (
	"github.com/gin-gonic/gin"
)

const (
	Ok         = 2000
	No         = 4000
	ServerFail = 5000
)

// OkWithData 成功的返回
func Success(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndData(c, Ok, message, data)
}

// FailWithData 客户端请求失败的返回
func Fail(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndData(c, No, message, data)
}

// ServerFailWithData 服务端响应失败的返回
func FailButServer(c *gin.Context, message string, data interface{}) {
	ResponseWithStatusAndData(c, ServerFail, message, data)
}

// ResponseWithStatusAndData 确定统一返回格式
func ResponseWithStatusAndData(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
