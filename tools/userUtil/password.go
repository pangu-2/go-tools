package userUtil

import (
	"github.com/pangu-2/go-tools/tools/cryptUtil"
)

// encrypt password
func PasswordSalt(pass, salt string) string {
	salt1 := "4%$@w"
	str := salt1 + pass + salt
	//return cryptUtil.Md5(cryptUtil.Sha256(str))
	return cryptUtil.Md5(str)
}

//验证
func PasswordVerify(password, pass, salt string) bool {
	return password == PasswordSalt(pass, salt)
}

//密码加密
func SaltMake(str, salt string) string {
	return cryptUtil.Md5(cryptUtil.Sha256(str+salt) + salt)
}
