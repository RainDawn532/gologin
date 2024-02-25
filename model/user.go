package model

import (
	"time"
)

type User struct {
	Id             int32
	Email          string
	Password       string
	Salt           string
	ConfirmCode    string
	ActivationTime time.Time
	IsValid        int8
}

func (User) TableName() string {
	return "t_user"
}
