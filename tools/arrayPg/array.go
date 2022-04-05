package arrayPg

import "reflect"

func IsArray(value interface{}) bool {
	rv := reflect.ValueOf(value)
	rk := rv.Kind()
	if rk == reflect.Ptr {
		rv = rv.Elem()
		rk = rv.Kind()
	}
	switch rk {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

func IsArrayGenerics(value any) bool {
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}
	rk := rv.Kind()
	if rk == reflect.Ptr {
		rv = rv.Elem()
		rk = rv.Kind()
	}
	switch rk {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}
