package convPg

import (
	"encoding/json"
	"github.com/pangu-2/go-tools/tools"
)

//StrToMap JSON格式数据转换为map
func StrToMap(str string) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	if str == "" {
		return nil, tools.NewError("字符串为空不能转换")
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		return nil, err
	}
	return result, nil
}
