package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string `json:"identity" gorm:"column:identity;type:varchar(36)"` // 分类的唯一标识
	Name     string `json:"name" gorm:"column:name;type:varchar(100)"`        // 分类ia
	ParentId string `json:"parent_id" gorm:"column:parent_id;type:int"`       // 父级id 默认为0
}

func (*Category) TableName() string {
	return "category"
}
