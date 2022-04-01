package cryptPg

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha1(v string) string {
	d := []byte(v)
	m := sha1.New()
	m.Write(d)
	//Sum()对hash.Hash对象内部存储的内容进行校验计算，然后将其追加到data的后面形成一个新的byte切片，所以一般需要将data设为nil
	//encoding/hex是一个将 byte切片转换成字符串的编码工具库
	return hex.EncodeToString(m.Sum(nil))
}

//func Sha12(msg string) string {
//	h := sha1.New()
//	io.WriteString(h, msg)
//	return fmt.Sprintf("%x", h.Sum(nil))
//}

//func Sha256(msg string) string {
//	h := sha256.New()
//	io.WriteString(h, msg)
//	return fmt.Sprintf("%x", h.Sum(nil))
//}

func Sha256(v string) string {
	d := []byte(v)
	m := sha256.New()
	m.Write(d)
	//Sum()对hash.Hash对象内部存储的内容进行校验计算，然后将其追加到data的后面形成一个新的byte切片，所以一般需要将data设为nil
	//encoding/hex是一个将 byte切片转换成字符串的编码工具库
	return hex.EncodeToString(m.Sum(nil))
}

func Sha512(v string) string {
	d := []byte(v)
	m := sha512.New()
	m.Write(d)
	//Sum()对hash.Hash对象内部存储的内容进行校验计算，然后将其追加到data的后面形成一个新的byte切片，所以一般需要将data设为nil
	//encoding/hex是一个将 byte切片转换成字符串的编码工具库
	return hex.EncodeToString(m.Sum(nil))
}

func Sha5122(v string) string {
	d := []byte(v)
	m := sha1.New()
	m.Write(d)
	//Sum()对hash.Hash对象内部存储的内容进行校验计算，然后将其追加到data的后面形成一个新的byte切片，所以一般需要将data设为nil
	//encoding/hex是一个将 byte切片转换成字符串的编码工具库
	return hex.EncodeToString(m.Sum(nil))
}
