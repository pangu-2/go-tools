package jsonPg

import (
	"github.com/pangu-2/go-tools/tools/convPg"
)

//ObjToJson obj 反序列化为字符串
func ObjToJson(v any) (string, error) {
	return convPg.ObjToJson(v)
}

//JsonEnCode 反序列化为字符串
func JsonEnCode(v interface{}) (string, error) {
	return convPg.JsonEnCode(v)
}
