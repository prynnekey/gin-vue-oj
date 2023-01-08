package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/models"
)

func GetUserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := models.GetUserList()
		if err != nil {
			ctx.JSON(http.StatusOK, "查询失败")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
