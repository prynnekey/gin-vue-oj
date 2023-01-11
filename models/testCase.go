package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	Identity        string `json:"identity"`         // 测试用例的唯一标识
	ProblemIdentity string `json:"problem_identity"` // 问题的唯一标识
	Input           string `json:"input"`            // 测试用例的输入
	Output          string `json:"output"`           // 测试用例的输出
}

func (*TestCase) TableName() string {
	return "test_case"
}
