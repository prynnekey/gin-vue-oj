package utils

import (
	"crypto/tls"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// 发送验证码
// 要发送的邮箱地址
// 发送的验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "在线联系系统 <prynnekey@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码"
	e.HTML = []byte("您的验证码为:<b>" + code + "</b>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "prynnekey@163.com", "EHOCMJSOCFNUCJBQ", "smtp.163.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})
	return err
}
