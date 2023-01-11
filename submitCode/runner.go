package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

// 执行用户提交代码的runner
func main() {
	// 执行用户的代码 go run userCode/main.go
	cmd := exec.Command("go", "run", "userCode/main.go")
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	// 进行标准输入
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdinPipe.Close()

	_, err = io.WriteString(stdinPipe, "11 20\n")
	if err != nil {
		fmt.Println(err)
	}

	// 读取标准输出
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	// 将读取到的内容与数据库答案进行对比
	fmt.Println(stdout.String())

	// 不相等

	// 相当

}
