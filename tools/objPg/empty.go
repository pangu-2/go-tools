package objPg

import (
	"reflect"

	"github.com/pangu-2/go-tools/tools/reflectionPg"
)

// IsEmpty
func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch result := value.(type) {
	case int:
		return result == 0
	case int8:
		return result == 0
	case int16:
		return result == 0
	case int32:
		return result == 0
	case int64:
		return result == 0
	case uint:
		return result == 0
	case uint8:
		return result == 0
	case uint16:
		return result == 0
	case uint32:
		return result == 0
	case uint64:
		return result == 0
	case float32:
		return result == 0
	case float64:
		return result == 0
	case bool:
		return result == false
	case string:
		return result == ""
	case []byte:
		//byte 等同于int8，常用来处理ascii字符
		return len(result) == 0
	case []rune:
		//rune 等同于int32,常用来处理unicode或utf-8字符
		return len(result) == 0
	case []int:
		return len(result) == 0
	case []string:
		return len(result) == 0
	case []float32:
		return len(result) == 0
	case []float64:
		return len(result) == 0
	case map[string]interface{}:
		return len(result) == 0

	default:
		var rv reflect.Value
		if v, ok := value.(reflect.Value); ok {
			rv = v
		} else {
			rv = reflect.ValueOf(value)
		}

		switch rv.Kind() {
		case reflect.Bool:
			return !rv.Bool()
		case
			reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64:
			return rv.Int() == 0

		case
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64,
			reflect.Uintptr:
			return rv.Uint() == 0

		case
			reflect.Float32,
			reflect.Float64:
			return rv.Float() == 0

		case reflect.String:
			return rv.Len() == 0

		case reflect.Struct:
			var fieldValueInterface interface{}
			for i := 0; i < rv.NumField(); i++ {
				fieldValueInterface, _ = reflectionPg.ValueToInterface(rv.Field(i))
				if !IsEmpty(fieldValueInterface) {
					return false
				}
			}
			return true

		case
			reflect.Chan,
			reflect.Map,
			reflect.Slice,
			reflect.Array:
			return rv.Len() == 0

		case
			reflect.Func,
			reflect.Ptr,
			reflect.Interface,
			reflect.UnsafePointer:
			if rv.IsNil() {
				return true
			}
		}
	}
	return false
}

// IsNil
func IsNil(value interface{}, traceSource ...bool) bool {
	if value == nil {
		return true
	}
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Map,
		reflect.Slice,
		reflect.Func,
		reflect.Interface,
		reflect.UnsafePointer:
		return !rv.IsValid() || rv.IsNil()

	case reflect.Ptr:
		if len(traceSource) > 0 && traceSource[0] {
			for rv.Kind() == reflect.Ptr {
				rv = rv.Elem()
			}
			if !rv.IsValid() {
				return true
			}
			if rv.Kind() == reflect.Ptr {
				return rv.IsNil()
			}
		} else {
			return !rv.IsValid() || rv.IsNil()
		}
	}
	return false
}
