package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/models"
)

// GetUserList 获取用户列表
// @Summary 获取所有用户
// @Schemes
// @Description 获取用户列表
// @Tags 公共方法
// @Accept json
// @Produce json
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /user-list [get]
func GetUserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := models.GetUserList()
		if err != nil {
			ctx.JSON(http.StatusOK, "查询失败")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	}
}
