package models

import "gorm.io/gorm"

type CategoryBasic struct {
	gorm.Model
	Identity string `json:"identity" gorm:"column:identity;type:varchar(36)"` // 分类的唯一标识
	Name     string `json:"name" gorm:"column:name;type:varchar(100)"`        // 分类ia
	ParentId string `json:"parent_id" gorm:"column:parent_id;type:int"`       // 父级id 默认为0
}

func (*CategoryBasic) TableName() string {
	return "category_basic"
}

func GetCategory(page, pageSize int, keyWord string) (*[]CategoryBasic, int64, error) {
	var categoryList *[]CategoryBasic
	var count int64
	err := DB.Model(&CategoryBasic{}).
		Count(&count).
		Where("name LIKE ?", "%"+keyWord+"%").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&categoryList).Error
	if err != nil {
		return nil, 0, err
	}

	return categoryList, count, nil
}
