package service

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/code"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/models"
)

// GetProblemList
// @Summary 获取问题列表
// @Param page query int false "请输入当前页,默认第一页"
// @Param pageSize query int false "每页多少条数据,默认20条"
// @Description 获取问题列表
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /problem-list [get]
func GetProblemList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
		pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))
		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, code.ERROR, "参数类型错误")
			return
		}

		proList, count, err := models.GetProblemList(page, pageSize)
		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, code.ERROR, "查询数据库失败")
			return
		}

		response.Success(ctx, code.OK, gin.H{
			"count": count,
			"list":  proList,
		}, "查询成功")
	}
}
