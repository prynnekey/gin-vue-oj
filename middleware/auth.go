package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/models"
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

		// 判断token是否过期
		redisToken, err := models.GetTokenWithRedis(uc.Username)
		if err != nil {
			response.Failed(ctx, "登录已过期,请重新登录")
			ctx.Abort()
			return
		}

		if redisToken != token {
			response.Failed(ctx, "token非法,请重新登陆")
			ctx.Abort()
			return
		}

		// 判断是否是管理员
		if uc == nil || uc.IsAdmin != 1 {
			// 不是 终止往下走
			response.Failed(ctx, "权限不足")
			ctx.Abort()
			return
		}

		// 是 放行
		ctx.Next()
	}
}

// 验证用户是否登录
func AuthUserCheck() gin.HandlerFunc {
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

		// 判断token是否过期
		redisToken, err := models.GetTokenWithRedis(uc.Username)
		if err != nil {
			response.Failed(ctx, "登录已过期,请重新登录")
			ctx.Abort()
			return
		}

		if redisToken != token {
			response.Failed(ctx, "token非法,请重新登陆")
			ctx.Abort()
			return
		}

		// 将用户id存储起来
		ctx.Set("user", uc)

		// 是 放行
		ctx.Next()
	}
}
