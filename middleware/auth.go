package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/utils"
)

// 权限验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从header中获取token
		token := ctx.GetHeader("Authorization")

		// 解析token
		uc, err := utils.ParseToken(token)
		if err != nil {
			response.Failed(ctx, "没有登录")
			ctx.Abort()
			return
		}

		// 判断是否是管理员
		if uc.IsAdmin != 1 {
			// 不是 终止往下走
			response.Failed(ctx, "权限不足")
			ctx.Abort()
			return
		}

		// 是 放行
		ctx.Next()
	}
}
