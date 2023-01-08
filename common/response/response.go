package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/code"
)

// 统一返回格式
func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// 成功
func Success(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code.OK,
		"data": data,
		"msg":  msg,
	})
}

// 失败
func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code.ERROR,
		"data": nil,
		"msg":  msg,
	})
}
