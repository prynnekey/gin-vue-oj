package service

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/define"
	"github.com/prynnekey/gin-vue-oj/models"
)

// GetCategoryList
// @Summary 获取分类列表
// @Param page query int false "请输入当前页,默认第一页"
// @Param pageSize query int false "每页多少条数据,默认20条"
// @Param keyWord query string false "关键字"
// @Description 获取分类列表
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /category/list [get]
func GetCategoryList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", define.PROBLEM_GET_PAGE))
		pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", define.PROBLEM_GET_PAGE_SIZE))
		keyWord := ctx.Query("keyWord")

		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, "参数类型错误")
			return
		}

		// 查询数据库
		cb, count, err := models.GetCategory(page, pageSize, keyWord)
		if err != nil {
			response.Failed(ctx, "查询失败")
			return
		}

		// 返回信息
		response.Success(ctx, gin.H{
			"count": count,
			"list":  cb,
		}, "查询成功")
	}
}
