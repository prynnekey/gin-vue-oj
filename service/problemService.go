package service

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/define"
	"github.com/prynnekey/gin-vue-oj/models"
	"gorm.io/gorm"
)

// GetProblemList
// @Summary 获取问题列表
// @Param page query int false "请输入当前页,默认第一页"
// @Param pageSize query int false "每页多少条数据,默认20条"
// @Param keyWord query string false "查询的关键字"
// @Param category_identity query string false "分类的唯一标识"
// @Description 获取问题列表
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /problem-list [get]
func GetProblemList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page, _ := strconv.Atoi(ctx.DefaultQuery("page", define.PROBLEM_GET_PAGE))
		pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", define.PROBLEM_GET_PAGE_SIZE))
		keyWord := ctx.Query("keyWord")
		categoryIdentity := ctx.Query("category_identity")

		if err != nil {
			log.Println("GetProblemList Param strconv Error:", err)
			response.Failed(ctx, "参数类型错误")
			return
		}

		proList, count, err := models.GetProblemList(page, pageSize, keyWord, categoryIdentity)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.Failed(ctx, "记录不存在")
				return
			}
			log.Println("GetProblemList Database Error:", err)
			response.Failed(ctx, "查询数据库失败")
			return
		}

		response.Success(ctx, gin.H{
			"count": count,
			"list":  proList,
		}, "查询成功")
	}
}

// GetProblemDetail
// @Summary 问题详情
// @Param identity query string false "问题的唯一标识"
// @Description 获取问题详细信息
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /problem-detail [get]
func GetProblemDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		identity := ctx.Query("identity")

		problem, err := models.GetProblemDetail(identity)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response.Failed(ctx, "数据不存在")
				return
			}
			response.Failed(ctx, "查询失败:"+err.Error())
			return
		}

		response.Success(ctx, problem, "查询成功")
	}
}

// AddProblem
// @Summary 添加一个问题
// @Description 添加问题
// @Tags 公共方法
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /problem-list [post]
func AddProblem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var problem models.ProblemBasic
		ctx.ShouldBindJSON(&problem)

		i, err := models.AddProblem(&problem)
		if err != nil {
			log.Println("AddProblem Error:", err)
			response.Failed(ctx, "err:"+err.Error())
			return
		}

		if i == 0 {
			response.Failed(ctx, "添加失败")
			return
		}

		response.Success(ctx, gin.H{
			"problem": problem,
		}, "添加成功")
	}
}
