package jsonPg

import (
	"encoding/json"
	"fmt"

	"github.com/pangu-2/go-tools/tools"
	"github.com/pangu-2/go-tools/tools/convPg"
)

//ObjToJson obj 反序列化为字符串
func ObjToJson(v any) (string, error) {
	return convPg.ObjToJson(v)
}

//JsonEnCode 反序列化为字符串
func JsonEnCode(v interface{}) (string, error) {
	str, err := json.Marshal(v)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return "", tools.NewError("序列化失败:" + err.Error())
	}
	return string(str), nil
}
