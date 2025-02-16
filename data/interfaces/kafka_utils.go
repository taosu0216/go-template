package interfaces

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

func (dr *DataRepo) SendVerifyCodeToEmailByKafka(verifyCode, to, subject string) error {
	e := email.Email{
		To:      []string{to},
		From:    dr.Cfg.MailCfg.From,
		HTML:    []byte(verifyCode),
		Subject: subject,
		Headers: textproto.MIMEHeader{},
	}

	auth := smtp.PlainAuth(
		"",
		dr.Cfg.MailCfg.From,     // 发件邮箱（需与From一致）
		dr.Cfg.MailCfg.Password, // IMAP/SMTP密码
		dr.Cfg.MailCfg.Host,     // SMTP服务器地址
	)
	err := e.SendWithTLS(
		dr.Cfg.MailCfg.Addr,
		auth,
		&tls.Config{
			ServerName:         dr.Cfg.MailCfg.Host, // 证书验证域名
			InsecureSkipVerify: true,                // 跳过证书验证（生产环境应设为false）
		},
	)

	if err != nil {
		return err
	}
	fmt.Println("邮件发送成功!")
	return nil
}
