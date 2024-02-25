package biz

import (
	"fmt"
	"testing"
)

func TestGenID(t *testing.T) {
	fmt.Println(GenID())
	fmt.Println(GenID())
	//var body bytes.Buffer

	//// 定义模板数据
	//data := struct {
	//	activationUrl string
	//}{
	//	activationUrl: "https://blog.csdn.net/qq_43514659/article/details/121362804",
	//}
	//file, _ := template.ParseFiles("D:\\goProject\\src\\gologin\\templates\\active-account.html")
	//
	//file.Execute(&body, "https://blog.csdn.net/qq_43514659/article/details/121362804")
	//fmt.Println(body.String())
}
