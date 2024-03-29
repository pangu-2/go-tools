package convPg

import (
	"encoding/json"

	"github.com/pangu-2/go-tools/tools"
)

//ObjToJson obj 反序列化为字符串
func ObjToJson(val any) (string, error) {
	str, err := json.Marshal(val)
	if err != nil {
		//fmt.Println("序列化失败:", err)
		return "", tools.NewError("序列化失败:" + err.Error())
	}
	return string(str), nil
}
