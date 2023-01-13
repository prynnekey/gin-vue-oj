package utils

import (
	"io/fs"
	"os"
	"strings"
)

// 将用户提交的代码保存到指定位置
// 保存的格式 submitCode/用户Identity/问题Identity/随机id/main.go
func SaveCode(userIdentity, problemIdentity string, code []byte) (string, error) {
	// 保存代码的路径
	dirName := "submitCode/" + userIdentity + "/" + problemIdentity + "/" + GenerateUUID()
	// 文件的名称
	path := dirName + "/main.go"

	// 创建文件夹
	// 777 全部权限
	err := os.MkdirAll(dirName, fs.ModePerm)
	if err != nil {
		return "", err
	}

	// 创建文件
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}

	// 将代码写入文件中
	_, err = f.Write(code)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return path, nil
}

func DeleteSaveCode(path string) error {
	s := strings.Split(path, "/")

	newPath := s[0] + "/" + s[1] + "/" + s[2] + "/" + s[3]

	err := os.RemoveAll(newPath)
	if err != nil {
		return err
	}

	return nil
}
