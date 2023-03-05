package strPg

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"math/rand"
	"time"
)

const DEFAULT_ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DEFAULT_ALPHABET_0 = "_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DEFAULT_Number = "0123456789"

// randseed
func GetRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// rand salt
func RandSalt() string {
	var salt = ""
	for i := 0; i < 4; i++ {
		ran := GetRand()
		salt += string(SALT[ran.Intn(len(SALT))])
	}
	return salt
}

const (
	SALT = "$^*#,.><)(_+f*m"
)

// 生成 数字和小写字母
// https://blog.csdn.net/qq948993066/article/details/77368971
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成随机字符串
func GetRandomString2(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetNanoid 生成随机字符串 gonanoid
func GetNanoid(l int) string {
	id := gonanoid.MustGenerate(DEFAULT_ALPHABET, l)
	return id
}

// GetNanoidGenerate 生成随机字符串 gonanoid
func GetNanoidGenerate(alphabet string, l int) string {
	id := gonanoid.MustGenerate(alphabet, l)
	return id
}

// GetNanoIdNumber 生成随机数值串
func GetNanoIdNumber(l int) string {
	id := gonanoid.MustGenerate(DEFAULT_Number, l)
	return id
}
