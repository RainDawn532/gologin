package handler

import (
	"github.com/gin-gonic/gin"
	"gologin/biz"
	"gologin/response"
)

func CreateAccount(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	flag, err := biz.CreateAccount(email, password)
	if flag {
		response.ResponseSuccessful(c, map[string]string{"message": "注册成功,请前往邮箱进行激活"})
	} else {
		response.ResponseWrong(c, 400, err.Error())
	}

}

func LoginAccount(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	flag, err, user := biz.LoginAccount(email, password)
	if flag {
		response.ResponseSuccessful(c, map[string]string{"message": "登录成功", "email": user.Email})
	} else {
		response.ResponseWrong(c, 400, err.Error())
	}
}

func ActivationAccount(c *gin.Context) {
	confirmCode, _ := c.GetQuery("confirmCode")
	flag, err := biz.ActivationAccount(confirmCode)
	//原始状态
	//if flag {
	//	c.JSON(200, gin.H{"message": "激活成功"})
	//} else {
	//	c.JSON(400, gin.H{"message": err.Error()})
	//}

	if flag {
		response.ResponseSuccessful(c, map[string]string{"message": "激活成功", "账号": "@lnh"})
	} else {
		response.ResponseWrong(c, 400, err.Error())
	}

}
