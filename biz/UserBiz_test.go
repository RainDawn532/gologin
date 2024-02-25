package biz

import (
	"gologin/internal"
	"testing"
)

func TestActivationAccount(t *testing.T) {
	internal.InitDB()
	flag, err := ActivationAccount("1761391492326363136")

	if err != nil || !flag {
		t.Error("Activation account failed,err", err)
	}

}

func TestCreateAccount(t *testing.T) {
	internal.InitDB()
	flag, err := CreateAccount("2071149217@qq.com", "qwer1234")

	if err != nil || !flag {
		t.Error("create account failed,err", err)
	}

}
