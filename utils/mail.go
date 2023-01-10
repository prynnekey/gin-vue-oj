package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/smtp"
	"strings"
	"time"

	"github.com/jordan-wright/email"
	"github.com/prynnekey/gin-vue-oj/define"
)

// 发送验证码
// 要发送的邮箱地址
// 发送的验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "在线联系系统 <prynnekey@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码"
	e.HTML = []byte("您的验证码为:<b>" + code + "</b>" + "五分钟内有效")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "prynnekey@163.com", "EHOCMJSOCFNUCJBQ", "smtp.163.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})
	return err
}

// 根据邮箱生成验证码
func GenerateCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	length := define.EMAIL_CODE_LENGTH

	var sb strings.Builder
	for i := 0; i < length; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}

	return sb.String()
}
