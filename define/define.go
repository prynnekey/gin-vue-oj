package define

var (
	PROBLEM_GET_PAGE      = "1"
	PROBLEM_GET_PAGE_SIZE = "20"

	// 用户注册时生成验证码的长度
	EMAIL_CODE_LENGTH = 6

	// 将验证码储存到Redis的格式
	REDIS_SAVE_EMAIL_CODE = "user:email_code:"
)
