package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"net/textproto"
)

func SendVerifyCodeToEmail(verifyCode, emailFrom, emailTo, addr, passwd, host, subject string) error {
	e := email.Email{
		To:      []string{emailTo},
		From:    emailFrom,
		HTML:    []byte(verifyCode),
		Subject: subject,
		Headers: textproto.MIMEHeader{},
	}

	auth := smtp.PlainAuth(
		"",
		emailFrom, // 发件邮箱（需与From一致）
		passwd,    // IMAP/SMTP密码
		host,      // SMTP服务器地址
	)
	err := e.SendWithTLS(
		addr,
		auth,
		&tls.Config{
			ServerName:         host, // 证书验证域名
			InsecureSkipVerify: true, // 跳过证书验证（生产环境应设为false）
		},
	)

	if err != nil {
		return err
	}
	fmt.Println("邮件发送成功!")
	return nil
}
