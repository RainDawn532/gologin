package biz

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
)

func GetMd5(s string) (string, error) {
	hash := md5.New()
	writeString, err := io.WriteString(hash, s)
	if writeString <= 0 {
		return "", errors.New("md5加密失败")
	}
	return hex.EncodeToString(hash.Sum(nil)), err
}
