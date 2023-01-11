package middleware

import "github.com/gin-gonic/gin"

// 权限验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token

		c.Next()
	}
}
