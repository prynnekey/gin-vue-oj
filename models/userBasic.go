package models

import (
	"context"
	"time"

	"github.com/prynnekey/gin-vue-oj/define"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity         string `json:"identity" gorm:"column:identity;type:varchar(36)"`             // 用户的唯一标识
	Username         string `json:"username" gorm:"column:username;type:varchar(100)"`            // 用户名
	Password         string `json:"password" gorm:"column:password;type:varchar(32)"`             // 密码
	Phone            string `json:"phone" gorm:"column:phone;type:varchar(20)"`                   // 手机号
	Mail             string `json:"mail" gorm:"column:mail;type:varchar(100)"`                    // 邮箱
	IsAdmin          int    `json:"is_admin" gorm:"column:is_admin;type:tinyint(1)"`              // 是否是管理员 1是 0不是 默认0
	FinishProblemNum int    `json:"finish_problem_num" gorm:"column:finish_problem_num;type:int"` // 完成的问题个数
	SubmitProblemNum int    `json:"submit_problem_num" gorm:"column:submit_problem_num;type:int"` // 提交的问题个数
}

func (*UserBasic) TableName() string {
	return "user_basic"
}

// 根据用户唯一标识获取用户
func GetUserDetail(identity string) (*UserBasic, error) {
	var user UserBasic
	err := DB.Where("identity = ?", identity).Omit("password").First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 根据用户名获取用户
func GetUserByUsername(username string) (*UserBasic, error) {
	var user UserBasic
	err := DB.Where("username = ?", username).Omit("password").First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 获取所有用户
func GetUserList() ([]UserBasic, error) {
	var user []UserBasic
	err := DB.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 获取用户排行榜
func GetUserRankList(page, pageSize int) (*[]UserBasic, int64, error) {
	var user []UserBasic
	var count int64
	// BUG: count统计不正确
	err := DB.Model(&UserBasic{}).
		Count(&count).
		Omit("identity", "password", "phone", "mail").
		Order("finish_problem_num DESC").
		Order("submit_problem_num").
		Order("created_at").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&user).Error
	if err != nil {
		return nil, 0, err
	}

	return &user, count, nil
}

func Login(username string) (*UserBasic, error) {
	var u UserBasic
	err := DB.Where("username = ?", username).First(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// 新增用户
func Register(userBasic *UserBasic) (int64, error) {
	d := DB.Create(&userBasic)
	return d.RowsAffected, d.Error
}

var ctx = context.Background()

// 将验证码存入Redis中
func SaveCodeWithRedis(mail, code string) error {
	// key   user:email:mail
	key := define.REDIS_SAVE_EMAIL_CODE + mail
	// 过期时间5分钟
	_, err := REDIS.SetNX(ctx, key, code, 5*time.Minute).Result()
	if err != nil {
		return err
	}

	return nil
}

// 获取Redis中的验证码
func GetCodeWithRedis(mail string) (string, error) {
	// key   user:email:mail
	key := define.REDIS_SAVE_EMAIL_CODE + mail
	// 过期时间5分钟
	code, err := REDIS.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return code, nil
}
