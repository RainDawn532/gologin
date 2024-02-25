package biz

import (
	"bytes"
	"fmt"
	"gologin/log"
	"html/template"
	"testing"
)

func TestSendActiveEmail(t *testing.T) {
	var body bytes.Buffer
	file, err := template.ParseFiles("D:\\goProject\\src\\gologin\\templates\\active-account.html")
	if err != nil {
		log.Logger.Warn("parse html failed" + err.Error())
	}

	file.Execute(&body, "https//:biying.com")

	fmt.Println(body.String())

	log.InitLogger()
	err2 := sendMailForActiveAccount("https://gorm.io/zh_CN/docs/create.html", "2071149217@qq.com")
	if err2 != nil {
		panic(err2)
	}
}
