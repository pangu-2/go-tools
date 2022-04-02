package convPg

import (
	"reflect"
	"strconv"

	"github.com/pangu-2/go-tools/generics"
)

//ObjToInt 变成数字
func ObjToInt(val interface{}) (ret int) {
	if nil == val {
		return 0
	}
	switch val.(type) {
	case int:
		ret = val.(int)
	case int8:
		ret = int(val.(int8))
	case int16:
		ret = int(val.(int16))
	case int32:
		ret = int(val.(int32))
	case int64:
		ret = int(val.(int64))
	case uint:
		ret = int(val.(uint))
	case uint8:
		ret = int(val.(uint8))
	case uint16:
		ret = int(val.(uint16))
	case uint32:
		ret = int(val.(uint32))
	case uint64:
		ret = int(val.(uint64))
	case uintptr:
		ret = int(val.(uintptr))
	case float32:
		ret = int(val.(float32))
	case float64:
		ret = int(val.(float64))
	case bool:
		if val.(bool) {
			ret = 1
		} else {
			ret = 0
		}
	case string:
		var err error
		ret, err = strconv.Atoi(val.(string))
		if err != nil {
			return 0
		}
	default:
		ret = 0
	}
	return ret
}

//ObjToInt32 变成数字
func ObjToInt32(val interface{}) (ret int32) {
	if nil == val {
		return 0
	}
	switch val.(type) {
	case int:
		ret = int32(val.(int))
	case int8:
		ret = int32(val.(int8))
	case int16:
		ret = int32(val.(int16))
	case int32:
		ret = val.(int32)
	case int64:
		ret = int32(val.(int64))

	case uint:
		ret = int32(val.(uint))
	case uint8:
		ret = int32(val.(uint8))
	case uint16:
		ret = int32(val.(uint16))
	case uint32:
		ret = int32(val.(uint32))
	case uint64:
		ret = int32(val.(uint64))
	case uintptr:
		ret = int32(val.(uintptr))
	case float32:
		ret = int32(val.(float32))
	case float64:
		ret = int32(val.(float64))
	case bool:
		if val.(bool) {
			ret = 1
		} else {
			ret = 0
		}
	case string:
		ret2, err := strconv.ParseInt(val.(string), 10, 32)
		if err != nil {
			ret = 0
		} else {
			ret = int32(ret2)
		}
	default:
		ret = 0
	}
	return ret
}

//ObjToInt64 变成数字
func ObjToInt64(val interface{}) (ret int64) {
	if nil == val {
		return 0
	}
	switch val.(type) {
	case int:
		ret = int64(val.(int))
	case int8:
		ret = int64(val.(int8))
	case int16:
		ret = int64(val.(int16))
	case int32:
		ret = int64(val.(int32))
	case int64:
		ret = val.(int64)

	case uint:
		ret = int64(val.(uint))
	case uint8:
		ret = int64(val.(uint8))
	case uint16:
		ret = int64(val.(uint16))
	case uint32:
		ret = int64(val.(uint32))
	case uint64:
		ret = int64(val.(uint64))
	case uintptr:
		ret = int64(val.(uintptr))
	case float32:
		ret = int64(val.(float32))
	case float64:
		ret = int64(val.(float64))
	case bool:
		if val.(bool) {
			ret = 1
		} else {
			ret = 0
		}
	case string:
		ret2, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			ret = 0
		} else {
			ret = int64(ret2)
		}
	default:
		ret = 0
	}
	return ret
}

//ObjToFloat32 变成数字
func ObjToFloat32(val interface{}) (ret float32) {
	if nil == val {
		return 0
	}
	switch val.(type) {
	case int:
		ret = float32(val.(int))
	case int8:
		ret = float32(val.(int8))
	case int16:
		ret = float32(val.(int16))
	case int32:
		ret = float32(val.(int32))
	case int64:
		ret = float32(val.(int64))

	case uint:
		ret = float32(val.(uint))
	case uint8:
		ret = float32(val.(uint8))
	case uint16:
		ret = float32(val.(uint16))
	case uint32:
		ret = float32(val.(uint32))
	case uint64:
		ret = float32(val.(uint64))
	case uintptr:
		ret = float32(val.(uintptr))
	case float32:
		ret = val.(float32)
	case float64:
		ret = float32(val.(float64))
	case bool:
		if val.(bool) {
			ret = 1
		} else {
			ret = 0
		}
	case string:
		ret2, err := strconv.ParseFloat(val.(string), 32)
		if err != nil {
			ret = 0
		} else {
			ret = float32(ret2)
		}
	default:
		ret = 0
	}
	return ret
}

//ObjToFloat64 变成数字
func ObjToFloat64(val interface{}) (ret float64) {
	if nil == val {
		return 0
	}
	switch val.(type) {
	case int:
		ret = float64(val.(int))
	case int8:
		ret = float64(val.(int8))
	case int16:
		ret = float64(val.(int16))
	case int32:
		ret = float64(val.(int32))
	case int64:
		ret = float64(val.(int64))

	case uint:
		ret = float64(val.(uint))
	case uint8:
		ret = float64(val.(uint8))
	case uint16:
		ret = float64(val.(uint16))
	case uint32:
		ret = float64(val.(uint32))
	case uint64:
		ret = float64(val.(uint64))
	case uintptr:
		ret = float64(val.(uintptr))
	case float32:
		ret = float64(val.(float32))
	case float64:
		ret = val.(float64)
	case bool:
		if val.(bool) {
			ret = 1
		} else {
			ret = 0
		}
	case string:
		var err error
		ret, err = strconv.ParseFloat(val.(string), 64)
		if err != nil {
			ret = 0
		}
	default:
		ret = 0
	}
	return ret
}

//ObjToInt onj变成数字
func ObjToInt2[T generics.Number](obj T) int {
	of := reflect.TypeOf(obj)
	ret := 0
	switch of.Kind() {
	case reflect.Int:
		ret = int(obj)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = int(obj)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ret = int(obj)
	case reflect.Float32, reflect.Float64:
		ret = int(obj)
	}
	return ret
}

//ObjToInt32 obj变成数字
func ObjToInt322[T generics.Number](obj T) int32 {
	of := reflect.TypeOf(obj)
	var ret int32 = 0
	switch of.Kind() {
	case reflect.Int:
		ret = int32(obj)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = int32(obj)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ret = int32(obj)
	case reflect.Float32, reflect.Float64:
		ret = int32(obj)
	}
	return ret
}

//ObjToInt64 obj变成数字
func ObjToInt642[T generics.Number](obj T) int64 {
	of := reflect.TypeOf(obj)
	var ret int64 = 0
	switch of.Kind() {
	case reflect.Int:
		ret = int64(obj)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = int64(obj)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ret = int64(obj)
	case reflect.Float32, reflect.Float64:
		ret = int64(obj)
	}
	return ret
}
