package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/prynnekey/gin-vue-oj/docs"
	"github.com/prynnekey/gin-vue-oj/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由规则

	// 问题
	r.GET("problem-list", service.GetProblemList())
	r.GET("problem-detail", service.GetProblemDetail())
	r.POST("problem-list", service.AddProblem())

	// 用户
	r.GET("user-detail", service.GetUserDetail())

	// 提交记录
	r.GET("submit-list", service.GetSubmitList())

	return r
}
