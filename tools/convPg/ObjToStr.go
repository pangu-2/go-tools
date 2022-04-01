package convPg

import (
	"strconv"
)

//ObjToStr obj 转换为 字符
func ObjToStr(val interface{}) (ret string) {
	if val == nil {
		return ""
	}
	switch val.(type) {
	case string:
		ret =  val.(string)
	case int:
		ret =  strconv.Itoa(val.(int))
	case int8:
		ret =  strconv.FormatInt(int64(val.(int8)), 10)
	case int16:
		ret =  strconv.FormatInt(int64(val.(int16)), 10)
	case int32:
		ret =  strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		ret =  strconv.FormatInt(val.(int64), 10)
	case uint:
		ret =  strconv.FormatUint(uint64(val.(uint)), 10)
	case uint8:
		ret =  strconv.FormatUint(uint64(val.(uint8)), 10)
	case uint16:
		ret =  strconv.FormatUint(uint64(val.(uint16)), 10)
	case uint32:
		ret =  strconv.FormatUint(uint64(val.(uint32)), 10)
	case uint64:
		ret =  strconv.FormatUint(val.(uint64), 10)
	case uintptr:
		ret =  strconv.FormatUint(uint64(val.(uintptr)), 10)
	case bool:
		ret =  strconv.FormatBool(val.(bool))
	case complex64:
		ret =  strconv.FormatComplex(complex128(val.(complex64)), 'f', 10, 128)
	case complex128:
		ret =  strconv.FormatComplex(val.(complex128), 'f', 6, 128)
	case float32:
		ret =  strconv.FormatFloat(float64(val.(float32)), 'f', -1, 64)
	case float64:
		ret =  strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case []byte:
		ret =  string(val.([]byte))
	default:
		ret, _ = ObjToJson(val)
	}
	return ret
}
