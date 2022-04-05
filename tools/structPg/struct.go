package structPg

import "reflect"

//IsStructType 类型
func IsStructType(value interface{}) bool {
	var (
		rv = reflect.ValueOf(value)
		rk  = rv.Kind()
	)
	for rk == reflect.Ptr {
		rv = rv.Elem()
		rk = rv.Kind()
	}
	switch rk {
	case reflect.Struct:
		return true
	}
	return false
}
