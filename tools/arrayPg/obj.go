package arrayPg

import (
	"encoding/json"
)

// ObjToMap
// 函　数：Obj2map
// 概　要：
// 参　数：
//      obj: 传入Obj
// 返回值：
//      mapObj: map对象
//      err: 错误
func ObjToMap(obj any) (mapObj map[string]interface{}, err error) {
	// 结构体转json
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &mapObj); err != nil {
		return nil, err
	}
	return mapObj, nil
}
