package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	// NOTE: 不明白为什么下面的gorm这么写
	Identity          string             `json:"identity" gorm:"column:identity;type:varchar(36)"`              // 问题的唯一标识
	ProblemCategories *[]ProblemCategory `json:"problem_categories" gorm:"foreignKey:problem_id;references:id"` // 关联问题的分类表
	Title             string             `json:"title" gorm:"column:title;type:varchar(255)"`                   // 问题的标题
	Content           string             `json:"content" gorm:"column:content;type:text"`                       // 问题的正文描述
	MaxMem            string             `json:"max_mem" gorm:"column:max_mem;type:int"`                        // 最大运行内存
	MaxRuntime        string             `json:"max_runtime" gorm:"column:max_runtime;type:int"`                // 最大运行时间
}

func (*ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(page int, pageSize int, keyWord string, categoryIdentity string) (*[]ProblemBasic, int64, error) {
	var problemList *[]ProblemBasic
	var count int64

	// 分页查询 查询第二页 每页10条
	// select * from problem limit 10 offset 10 orderby update_at
	// BUG: count计数与实际不符合
	tx := DB.Model(&ProblemBasic{}).
		Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ?", "%"+keyWord+"%", "%"+keyWord+"%").
		Count(&count)

	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	// NOTE: 下面代码看不懂
	if categoryIdentity != "" {
		err := tx.Joins("RIGHT JOIN problem_category pc ON pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity).Error
		if err != nil {
			return nil, 0, err
		}
	}

	err := tx.
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&problemList).Error

	if err != nil {
		return nil, 0, err
	}

	return problemList, count, nil
}

func GetProblemDetail(identity string) (*ProblemBasic, error) {
	var problem ProblemBasic

	err := DB.Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		Where("identity = ?", identity).
		First(&problem).Error
	if err != nil {
		return nil, err
	}

	return &problem, nil
}

// 添加一条数据 返回影响的行数和错误信息
func AddProblem(pro *ProblemBasic) (int64, error) {
	d := DB.Model(&ProblemBasic{}).Create(&pro)
	return d.RowsAffected, d.Error
}
