package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string        `json:"identity" gorm:"column:identity;type:varchar(36)"`                     // 分类的唯一标识
	ProblemIdentity string        `json:"problem_identity" gorm:"column:problem_identity;type:varchar(36)"`     // 问题的唯一标识
	ProblemBasic    *ProblemBasic `json:"problem_basic" gorm:"foreignKey:identity;references:problem_identity"` // 关联问题基础表
	UserIdentity    string        `json:"user_identity" gorm:"column:user_identity;type:varchar(36)"`           // 用户的唯一表四
	UserBasic       *UserBasic    `json:"user_basic" gorm:"foreignKey:identity;references:user_identity"`       // 关联用户基础表
	CodePath        string        `json:"code_path" gorm:"column:code_path;type:varchar(255)"`                  // 代码存放路径
	Status          int           `json:"status" gorm:"column:status;type:tinyint(1)"`                          // 状态码 【-1-待判断,1-答案正确,2-答案错误,3-运行超时,4-运行超内存】
}

func (*SubmitBasic) TableName() string {
	return "submit_basic"
}

// 获取提交列表
func GetSubmitList(page, pageSize int, problemIdentity, userIdentity string, status int) (*[]SubmitBasic, int64, error) {
	tx := DB.Model(&SubmitBasic{}).
		Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
			// 忽略content字段
			return db.Omit("content")
		}).
		Preload("UserBasic")

	// 条件查询
	if problemIdentity != "" {
		err := tx.Where("problem_identity = ?", problemIdentity).Error
		if err != nil {
			return nil, 0, err
		}
	}

	if userIdentity != "" {
		err := tx.Where("user_identity = ?", userIdentity).Error
		if err != nil {
			return nil, 0, err
		}
	}

	if status != 0 {
		err := tx.Where("status = ?", status).Error
		if err != nil {
			return nil, 0, err
		}
	}

	var count int64
	var submitList *[]SubmitBasic

	// BUG: count总是不对
	// 分页查询
	err := tx.Count(&count).Limit(pageSize).Offset((page - 1) * pageSize).Find(&submitList).Error
	if err != nil {
		return nil, 0, err
	}

	return submitList, count, nil
}
