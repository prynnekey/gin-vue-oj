package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `json:"identity" gorm:"column:identity;type:varchar(36)"`  // 用户的唯一标识
	Username string `json:"username" gorm:"column:username;type:varchar(100)"` // 用户名
	Password string `json:"password" gorm:"column:password;type:varchar(32)"`  // 密码
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(20)"`        // 手机号
	Mail     string `json:"mail" gorm:"column:mail;type:varchar(100)"`         // 邮箱
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

// 获取所有用户
func GetUserList() ([]UserBasic, error) {
	var user []UserBasic
	err := DB.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(username string) (*UserBasic, error) {
	var u UserBasic
	err := DB.Where("username = ?", username).First(&u).Error
	if err != nil {
		return nil, err
	}

	return &u, nil
}
