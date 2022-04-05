package numberPg

//IsFloatGt0 检测 是否 大于0
func IsFloatGt0(val float64) bool {
	return 0 < val
}

//IsFloatGt320 检测 是否 大于0
func IsFloatGt320(val float32) bool {
	return 0 < val
}

//IsFloatGt640 检测 是否 大于0
func IsFloatGt640(val float64) bool {
	return 0 < val
}

//IsFloatType 类型
func IsFloatType(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}
