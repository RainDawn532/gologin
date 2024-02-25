package biz_test

import (
	"fmt"
	"gologin/biz"
	"testing"
)

func TestGetMd5(t *testing.T) {
	s := "qwer1234" + "wkb223"
	fmt.Println(biz.GetMd5(s))
	s1 := "test2"
	fmt.Println(biz.GetMd5(s1))
}
