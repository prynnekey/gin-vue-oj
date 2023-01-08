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
	r.GET("test", service.Ping())
	r.GET("user-list", service.GetUserList())
	r.GET("problem-list", service.GetProblemList())

	return r
}
