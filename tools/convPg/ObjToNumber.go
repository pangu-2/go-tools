package convPg

import (
	"github.com/pangu-2/go-tools/generics"
	"reflect"
)

//ObjToInt onj变成数字
func ObjToInt[T generics.Number](obj T) int {
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
func ObjToInt32[T generics.Number](obj T) int32 {
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
func ObjToInt64[T generics.Number](obj T) int64 {
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
