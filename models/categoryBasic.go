package models

import (
	"github.com/prynnekey/gin-vue-oj/utils"
	"gorm.io/gorm"
)

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

func AddCategory(name, parentId string) error {
	err := DB.Create(&CategoryBasic{
		Identity: utils.GenerateUUID(),
		Name:     name,
		ParentId: parentId,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

// 根据id删除分类 如果该分类下有数据 则不允许删除
// 返回类型分别是 删除的行数、该分类下是否有数据(有数据true) 和 错误信息
func DeleteCategoryById(id string) (int64, bool, error) {
	var count int64
	err := DB.Model(&ProblemCategory{}).Where("category_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, true, err
	}

	if count > 0 {
		// 说明该分类下有数据 不能删除
		return 0, true, nil
	}

	tx := DB.Where("id = ?", id).Delete(&CategoryBasic{})
	return tx.RowsAffected, false, tx.Error
}

func UpdateCategoryById(id, name, parentId string) (int64, error) {
	tx := DB.Model(&CategoryBasic{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":      name,
		"parent_id": parentId,
	})
	return tx.RowsAffected, tx.Error
}
