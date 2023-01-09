package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/models"
	"gorm.io/gorm"
)

// GetUserDetail
// @Summary 获取用户详细信息
// @Description 获取用户详细信息
// @Tags 公共方法
// @Param identity query string false "用户的唯一标识"
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /user-detail [get]
func GetUserDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		identity := ctx.Query("identity")

		// 参数校验
		if identity == "" {
			response.Failed(ctx, "用户唯一标识不能为空")
			return
		}

		// 查询数据库
		u, err := models.GetUserDetail(identity)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.Failed(ctx, "该用户不存在")
				return
			}
			response.Failed(ctx, "err:"+err.Error())
			return
		}

		// 返回结果
		response.Success(ctx, u, "查询成功")
	}
}

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
