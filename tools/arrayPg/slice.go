package arrayPg

import "reflect"

// SliceContainsString 判断是否包含
func SliceContainsString(src []string, val string) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == val {
			isContain = true
			break
		}
	}
	return isContain
}

// SliceContainsInt 判断是否包含
func SliceContainsInt(src []int, value int) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

// SliceContainsInt64 判断是否包含
func SliceContainsInt64(src []int64, value int64) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

// SliceContainsInt32 判断是否包含
func SliceContainsInt32(src []int32, value int32) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

// SliceContainsInt16 判断是否包含
func SliceContainsInt16(src []int16, value int16) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

// SliceContainsInt8 判断是否包含
func SliceContainsInt8(src []int8, value int8) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
			isContain = true
			break
		}
	}
	return isContain
}

//IsSliceType 类型
func IsSliceType(val interface{}) bool {
	var (
		rv = reflect.ValueOf(val)
		rk = rv.Kind()
	)
	for rk == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rk {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}
