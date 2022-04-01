package cryptPg

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	//Sum()对hash.Hash对象内部存储的内容进行校验计算，然后将其追加到data的后面形成一个新的byte切片，所以一般需要将data设为nil
	//encoding/hex是一个将 byte切片转换成字符串的编码工具库
	return hex.EncodeToString(m.Sum(nil))
}

func Md52(msg string) string {
	h := md5.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}
