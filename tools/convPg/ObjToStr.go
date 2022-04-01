package convPg

import (
	"reflect"
	"strconv"
)

//ObjToStr obj 转换为 字符
func ObjToStr(i interface{}) string {
	if i == nil {
		return ""
	}
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int8:
		return strconv.FormatInt(int64(i.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(i.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(i.(int32)), 10)
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case uint:
		return strconv.FormatUint(uint64(i.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(i.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(i.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(i.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(i.(uint64), 10)
	case uintptr:
		return strconv.FormatUint(uint64(i.(uintptr)), 10)
	case bool:
		return strconv.FormatBool(i.(bool))
	case complex64:
		return strconv.FormatComplex(complex128(i.(complex64)), 'f', 10, 128)
	case complex128:
		return strconv.FormatComplex(i.(complex128), 'f', 6, 128)
	case float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', 6, 64)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', 6, 64)
	}
	k := reflect.TypeOf(i).Kind()
	if k == reflect.Struct || reflect.Array == k || reflect.Map == k || reflect.Slice == k {
		v, err := ObjToJson(i)
		if err == nil {
			return v
		}
		return ""
	}

	return ""
}
