package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `json:"identity" gorm:"column:identity;type:varchar(36)"`                 // 分类的唯一标识
	ProblemIdentity string `json:"problem_identity" gorm:"column:problem_identity;type:varchar(36)"` // 问题的唯一标识
	UserIdentity    string `json:"user_identity" gorm:"column:user_identity;type:varchar(36)"`       // 用户的唯一表四
	CodePath        string `json:"code_path" gorm:"column:code_path;type:varchar(255)"`              // 代码存放路径
	Status          int    `json:"status" gorm:"column:status;type:tinyint"`                         // 状态码 【0-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存】
}

func (*Submit) TableName() string {
	return "submit"
}
