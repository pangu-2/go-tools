package arrayPg

// SliceContainsString 判断是否包含
func SliceContainsString(src []string, value string) bool {
	isContain := false
	for _, srcValue := range src {
		if srcValue == value {
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
