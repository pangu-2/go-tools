package numberPg

import "strconv"

// IsFloatGt0 检测 是否 大于0
func IsFloatGt0(val float64) bool {
	return 0 < val
}

// IsFloatGt320 检测 是否 大于0
func IsFloatGt320(val float32) bool {
	return 0 < val
}

// IsFloatGt640 检测 是否 大于0
func IsFloatGt640(val float64) bool {
	return 0 < val
}

// IsFloatType 类型
func IsFloatType(value interface{}) bool {
	switch value.(type) {
	case float32, *float32, float64, *float64:
		return true
	}
	return false
}

// Float64ToStr
//
//	@Description: float64 转换为字符
//	@param val
//	@return string
func Float64ToStr(val float64) string {
	return strconv.FormatFloat(val, 'E', -1, 64)
}

// Float64ToStrPrec
//
//	@Description: float64 转换为字符，有精度
//	@param val
//	@param precision
//	@return string
func Float64ToStrPrec(val float64, precision int) string {
	return strconv.FormatFloat(val, 'E', precision, 64)
}
