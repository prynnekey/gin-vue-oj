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
		password = utils.MD5(password)

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
		code := utils.GenerateCode()

		// 将生成的验证码存入Redis中
		err := models.SaveCodeWithRedis(email, code)
		if err != nil {
			response.Failed(ctx, "发生错误"+err.Error())
			return
		}

		// 发送邮箱
		err = utils.SendCode(email, code)
		if err != nil {
			response.Failed(ctx, "发送失败:"+err.Error())
			return
		}

		response.Success(ctx, code, "验证码发送成功")
	}
}

// Register
// @Summary 用户注册
// @Description 用户注册
// @Tags 公共方法
// @Param username formData string false "用户名"
// @Param password formData string false "密码"
// @Param confirm_password formData string false "确认密码"
// @Param phone formData string false "手机号"
// @Param mail formData string false "邮箱"
// @Param code formData string false "验证码"
// @Success 200 {string} json "{“code”: "200", "msg":"", "data": ""}"
// @Router /register [post]
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		confirmPassword := ctx.PostForm("confirm_password")
		phone := ctx.PostForm("phone")
		mail := ctx.PostForm("mail")
		code := ctx.PostForm("code")

		// 校验数据格式
		if username == "" {
			response.Failed(ctx, "用户名为空")
			return
		}

		if password == "" {
			response.Failed(ctx, "密码为空")
			return
		}

		if confirmPassword == "" {
			response.Failed(ctx, "确认密码为空")
			return
		}

		if phone == "" {
			response.Failed(ctx, "手机号为空")
			return
		}

		if mail == "" {
			response.Failed(ctx, "邮箱为空")
			return
		}

		if code == "" {
			response.Failed(ctx, "验证码为空")
			return
		}

		if password != confirmPassword {
			response.Failed(ctx, "两次密码不一致")
			return
		}

		// 如果用户名存在 则返回错误信息
		_, err := models.GetUserByUsername(username)
		if err == nil {
			response.Failed(ctx, "用户名已存在")
			return
		}

		if err != gorm.ErrRecordNotFound {
			response.Failed(ctx, "用户名已存在")
			return
		}

		// 根据邮箱从Redis中获取验证码
		redisCode, _ := models.GetCodeWithRedis(mail)

		// 对比输入的验证码是否正确
		if code != redisCode {
			// 不正确
			response.Failed(ctx, "验证码不正确")
			return
		}

		// 将密码进行md5加密
		password = utils.MD5(password)

		// 生成用户唯一标识 uuid
		identity := utils.GenerateUUID()

		// 将数据封装为user类
		user := models.UserBasic{
			Identity: identity,
			Username: username,
			Password: password,
			Phone:    phone,
			Mail:     mail,
		}

		// 将数据插入数据库
		i, err := models.Register(&user)
		if err != nil || i == 0 {
			response.Failed(ctx, "注册失败:"+err.Error())
			return
		}

		// 注册成功 生成token 使得用户直接登录
		token, err := utils.GenerateToken(identity, username)
		if err != nil {
			response.Failed(ctx, "生成token失败，请重新登陆:"+err.Error())
			return
		}
		response.Success(ctx, gin.H{
			"token": token,
		}, "注册成功")
	}
}
