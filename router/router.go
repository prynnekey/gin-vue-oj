package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/service"
)

func Init() *gin.Engine {
	r := gin.Default()

	// 路由规则
	r.GET("test", service.Ping())
	r.GET("user-list", service.GetUserList())

	return r
}
