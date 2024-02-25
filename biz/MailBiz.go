package biz

import (
	"bytes"
	"gologin/log"
	"html/template"
	"net/smtp"
	"time"
)

// 网易邮箱 smtp 服务器地址
const host = "smtp.163.com"

// 端口
const port = "25"

// 发件人邮箱地址
const username = "lnh611@163.com"

// 授权码
const password = "ODLEZQKPRAVIZPFK"

func sendMailForActiveAccount(activationUrl string, email string) error {

	var body bytes.Buffer
	file, err := template.ParseFiles("D:\\goProject\\src\\gologin\\templates\\active-account.html")
	if err != nil {
		log.Logger.Warn("parse html failed" + err.Error())
	}

	file.Execute(&body, activationUrl)
	// 构建邮件体
	message := "To: " + email + "\r\n" +
		"Subject: " + "账号激活" + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		body.String()

	// 连接SMTP服务器
	//503 认证失败  最主要是password其实是授权码  而不是密码
	auth := smtp.PlainAuth("", username, password, host)
	err = smtp.SendMail(host+":"+port, auth, username, []string{email}, []byte(message))
	if err != nil {
		log.Logger.Info(err.Error())
		return err
	}

	log.Logger.Info(email + "邮件发送成功 " + "时间" + time.Now().String())

	return nil
}
