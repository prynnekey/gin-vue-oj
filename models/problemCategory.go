package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     string         `json:"problem_id" gorm:"column:problem_id"`                        // 问题的id
	CategoryId    string         `json:"category_id" gorm:"column:category_id"`                      // 分类的id
	CategoryBasic *CategoryBasic `json:"category_basic" gorm:"foreignKey:id;references:category_id"` // 关联分类的基础信息表
}

func (*ProblemCategory) TableName() string {
	return "problem_category"
}
