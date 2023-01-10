package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/gin-vue-oj/common/response"
	"github.com/prynnekey/gin-vue-oj/models"
	"github.com/prynnekey/gin-vue-oj/utils"
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

// Login
// @Summary 用户登录
// @Description 用户登录
// @Tags 公共方法
// @Param username formData string false "用户名"
// @Param password formData string false "密码"
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /login [post]
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取登录的用户名和密码
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		// 校验输入格式
		if username == "" || password == "" {
			response.Failed(ctx, "用户名或密码不能为空")
			return
		}

		// 将密码进行md5加密
		// password = utils.MD5(password)

		// 根据用户名查询数据
		ub, err := models.Login(username)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// 没查到 提示用户不存在
				response.Failed(ctx, "用户不存在")
				return
			}
			response.Failed(ctx, "err:"+err.Error())
			return
		}

		// 将查询到的数据与用户输入的数据对比
		if ub.Password != password {
			// 对比失败 返回密码错误
			response.Failed(ctx, "密码错误")
			return
		}

		// 对比成功 生成token
		tokenString, err := utils.GenerateToken(ub.Identity, ub.Username)
		if err != nil {
			response.Failed(ctx, "生成token失败:"+err.Error())
			return
		}

		response.Success(ctx, gin.H{
			"token": tokenString,
		}, "登录成功")
	}
}

// SendCode
// @Summary 发送邮箱验证码
// @Description 发送邮箱验证码
// @Tags 公共方法
// @Param email formData string false "用户邮箱"
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /send-code [post]
func SendCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取邮箱
		email := ctx.PostForm("email")

		// 校验邮箱格式
		if email == "" {
			response.Failed(ctx, "输入的电子邮箱为空")
			return
		}

		// 生成验证码
		code := "123456"

		// 发送邮箱
		err := utils.SendCode(email, code)
		if err != nil {
			response.Failed(ctx, "发送失败:"+err.Error())
			return
		}

		response.Success(ctx, code, "验证码发送成功")
	}
}
