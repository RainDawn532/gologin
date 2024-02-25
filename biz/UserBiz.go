package biz

import (
	"errors"
	"gologin/internal"
	"gologin/log"
	"gologin/model"
	"gologin/util"
	"strconv"
	"time"
)

func CreateAccount(email string, password string) (bool, error) {
	var user model.User
	userResult := internal.DB.Where("email = ? ", email).Find(&user)
	if userResult.RowsAffected > 0 {
		return false, errors.New("账户已存在,请前往激活或者登录")
	}

	//雪花算法确认码
	confirmCode := GenID()
	salt := util.RandomString(6)
	passwordMd5, err := GetMd5(password + salt)
	if err != nil {
		log.Logger.Info("get md5 password fail" + err.Error())
	}
	//激活到期时间
	ldt := time.Now().Add(time.Hour * 24)
	user.Email = email
	user.Salt = salt
	user.ConfirmCode = strconv.Itoa(int(confirmCode))
	user.Password = passwordMd5
	user.IsValid = 0
	user.ActivationTime = ldt
	result := internal.DB.Create(&user)
	if result.RowsAffected <= 0 {
		return false, errors.New("插入失败")
	}

	//进一步优化异步进行发送
	activationUrl := "http://localhost:8080/activation?confirmCode=" + user.ConfirmCode
	go func() {
		err := sendMailForActiveAccount(activationUrl, user.Email)
		if err != nil {
			log.Logger.Warn("发送邮件失败,联系管理员")
		}
	}()

	return true, nil
}

func LoginAccount(email string, password string) (bool, error, model.User) {
	var user []model.User
	result := internal.DB.Where("email = ? AND is_valid = 1", email).Find(&user)
	if result.RowsAffected <= 0 {
		return false, errors.New("不存在"), model.User{}
	}
	u := user[0]
	md5Password, err := GetMd5(password + u.Salt)
	if nil == err {
		if len(user) == 1 {
			if u.Password == md5Password {
				return true, nil, u
			} else {
				return false, errors.New("账号密码错误"), u
			}
		}
	}
	return false, errors.New("账户异常,请点击右下角联系管理员"), model.User{}
}

func ActivationAccount(confirmCode string) (bool, error) {
	var user model.User
	result := internal.DB.Where("confirm_code = ? AND is_valid = 0", confirmCode).First(&user)
	if result.RowsAffected <= 0 {
		return false, errors.New("账号已经激活")
	}
	if time.Now().After(user.ActivationTime) {
		return false, errors.New("已失效请重新注册")
	}
	//db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;
	result1 := internal.DB.Model(&user).Where("confirm_code = ? ", confirmCode).Update("is_valid", 1)
	if result1.RowsAffected <= 0 {
		return false, errors.New("激活失败")
	}
	return true, nil
}
