package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `json:"identity" gorm:"column:identity;type:varchar(36)"`  // 用户的唯一标识
	Username string `json:"username" gorm:"column:username;type:varchar(100)"` // 用户名
	Password string `json:"password" gorm:"column:password;type:varchar(32)"`  // 密码
	Phone    string `json:"phone" gorm:"column:phone;type:varchar(20)"`        // 手机号
	Mail     string `json:"mail" gorm:"column:mail;type:varchar(100)"`         // 邮箱
}

func (*User) TableName() string {
	return "user"
}

// 获取所有用户
func GetUserList() ([]User, error) {
	var user []User
	err := DB.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}