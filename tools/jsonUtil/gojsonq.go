package jsonUtil

import (
	"encoding/json"
	"github.com/thedevsaddam/gojsonq/v2"
)

//根据json 字符串,取字段数据
func GetJson(jsonStr, field string) interface{} {
	return gojsonq.New().FromString(jsonStr).Find(field)
}

//根据json 字符串,取字段数据
func GetJsonString(jsonStr, field string) string {
	return gojsonq.New().FromString(jsonStr).Find(field).(string)
}

//根据json 字符串,取字段数据
func GetJsonInt(jsonStr, field string) int {
	return gojsonq.New().FromString(jsonStr).Find(field).(int)
}

//根据json 字符串,取字段数据
func GetJsonInt32(jsonStr, field string) int32 {
	return gojsonq.New().FromString(jsonStr).Find(field).(int32)
}

//根据json 字符串,取字段数据
func GetJsonInt64(jsonStr, field string) int64 {
	return gojsonq.New().FromString(jsonStr).Find(field).(int64)
}

//根据json 字符串,取字段数据
func GetJsonBool(jsonStr, field string) bool {
	return gojsonq.New().FromString(jsonStr).Find(field).(bool)
}

//根据json 字符串,取字段数据
func GetJsonIntArr(jsonStr, field string) []int {
	return gojsonq.New().FromString(jsonStr).Find(field).([]int)
}

//根据json 字符串,取字段数据
func GetJsonInt32Arr(jsonStr, field string) []int32 {
	return gojsonq.New().FromString(jsonStr).Find(field).([]int32)
}

//根据json 字符串,取字段数据
func GetJsonInt64Arr(jsonStr, field string) []int64 {
	return gojsonq.New().FromString(jsonStr).Find(field).([]int64)
}

//根据json 字符串,取字段数据
func GetJsonStringArr(jsonStr, field string) []string {
	return gojsonq.New().FromString(jsonStr).Find(field).([]string)
}

//根据Struct,取字段数据
func GetJsonByStruct(jsonS interface{}, field string) (interface{}, error) {
	jsonStr, err := json.Marshal(jsonS)
	if err != nil {
		return nil, err
	}
	return gojsonq.New().FromString(string(jsonStr)).Find(field), nil
}
