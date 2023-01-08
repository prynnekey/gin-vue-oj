package service

import "github.com/gin-gonic/gin"

func Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ping": "pang",
		})
	}
}
