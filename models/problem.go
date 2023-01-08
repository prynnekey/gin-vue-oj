package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `json:"identity" gorm:"column:identity;type:varchar(36)"`        // 问题的唯一标识
	CategoryId string `json:"category_id" gorm:"column:category_id;type:varchar(255)"` // 分类id
	Title      string `json:"title" gorm:"column:title;type:varchar(255)"`             // 问题的标题
	Content    string `json:"content" gorm:"column:content;type:text"`                 // 问题的正文描述
	MaxMem     string `json:"max_mem" gorm:"column:max_mem;type:int"`                  // 最大运行内存
	MaxRuntime string `json:"max_runtime" gorm:"column:max_runtime;type:int"`          // 最大运行时间
}

func (*Problem) TableName() string {
	return "problem"
}
