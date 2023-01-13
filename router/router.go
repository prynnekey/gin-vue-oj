package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/prynnekey/gin-vue-oj/docs"
	"github.com/prynnekey/gin-vue-oj/middleware"
	"github.com/prynnekey/gin-vue-oj/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 路由规则

	// 公有方法
	// 问题
	problem := r.Group("/problem")
	{
		problem.GET("/list", service.GetProblemList())
		problem.GET("/detail", service.GetProblemDetail())
	}

	// 用户
	user := r.Group("/user")
	{
		user.GET("/detail", service.GetUserDetail())
		user.POST("/login", service.Login())
		user.POST("/register", service.Register())
		user.POST("/send-code", service.SendCode())

		// 用户提交代码的方法 需要验证用户是否登录
		user.POST("/submit", middleware.AuthUserCheck(), service.SubmitCode())
	}

	// 排行榜
	r.GET("/rank-list", service.GetRankList())

	// 提交记录
	r.GET("submit-list", service.GetSubmitList())

	// 管理员私有方法
	admin := r.Group("/admin", middleware.AuthMiddleware())
	{
		// 新增问题
		admin.POST("/problem", service.AddProblem())
		// 问题修改
		admin.PUT("/problem", service.UpdateProblemByIdentity())

		// 查看所有用户
		admin.GET("/user-list", service.GetUserList())

		// 分类
		category := admin.Group("/category")
		{
			// 获取分类列表
			category.GET("/", service.GetCategoryList())
			category.POST("/", service.AddCategory())
			category.DELETE("", service.DeleteCategoryById())
			category.PUT("", service.UpdateCategoryById())
		}
	}

	return r
}
