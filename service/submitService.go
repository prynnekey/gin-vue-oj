package service

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/define"
	"github.com/prynnekey/gin-vue-oj/models"
)

// GetSubmitList
// @Summary 获取提交记录列表
// @Param page query int false "请输入当前页,默认第一页"
// @Param pageSize query int false "每页多少条数据,默认20条"
// @Param problem_identity query string false "问题的唯一标识"
// @Param user_identity query string false "用户的唯一标识"
// @Param status query int false "提交的状态【-1-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存】"
// @Description 获取问题列表
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /submit-list [get]
func GetSubmitList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", define.PROBLEM_GET_PAGE))
		pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", define.PROBLEM_GET_PAGE_SIZE))
		problemIdentity := ctx.Query("problem_identity")
		userIdentity := ctx.Query("user_identity")
		status, _ := strconv.Atoi(ctx.Query("status"))

		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, "参数类型错误")
			return
		}

		// 查询数据库
		submitList, count, err := models.GetSubmitList(page, pageSize, problemIdentity, userIdentity, status)
		if err != nil {
			response.Failed(ctx, "查询数据库出错:"+err.Error())
			return
		}

		response.Success(ctx, gin.H{"count": count, "list": submitList}, "查询成功")
	}
}
