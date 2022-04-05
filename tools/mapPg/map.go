package mapPg

import "reflect"

// MapContains 判断key是否存在
func MapContains(src map[string]interface{}, key string) bool {
	if _, ok := src[key]; ok {
		return true
	}
	return false
}

//IsMapType 类型
func IsMapType(value interface{}) bool {
	var (
		rv = reflect.ValueOf(value)
		rk = rv.Kind()
	)
	for rk == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rk {
	case reflect.Map:
		return true
	}
	return false
}
