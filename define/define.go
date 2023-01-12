package define

import "time"

var (
	PROBLEM_GET_PAGE      = "1"
	PROBLEM_GET_PAGE_SIZE = "20"

	// 用户注册时生成验证码的长度
	EMAIL_CODE_LENGTH = 6

	// 将验证码储存到Redis的格式
	REDIS_SAVE_EMAIL_CODE = "user:email_code:"
	// 验证码过期时间 5分钟
	REDIS_SAVE_EMAIL_CODE_EXPIRY = 5 * time.Minute

	// token存入Redis的格式
	REDIS_SAVE_USER_TOKEN = "user:token:"
	// token的过期时间 7天
	REDIS_SAVE_USER_TOKEN_EXPIRY = 7 * 24 * time.Hour
)
