package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5 str 取MD5
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}
