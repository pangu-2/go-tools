package convPg

import "strconv"

// Float32tToStr
//
//	@Description: float 转换为 字符
//	@param val
//	@return string
func Float32tToStr(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', -1, 64)
}

// Float32tToStrPrec
//
//	@Description: float 转换为 字符，有精度
//	@param val
//	@param precision
//	@return string
func Float32tToStrPrec(val float32, precision int) string {
	return strconv.FormatFloat(float64(val), 'f', precision, 64)
}

// Float64ToStr
//
//	@Description: float 转换为 字符
//	@param val
//	@return string
func Float64ToStr(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

// Float64ToStrPrec
//
//	@Description: float 转换为 字符，有精度
//	@param val
//	@param precision
//	@return string
func Float64ToStrPrec(val float64, precision int) string {
	return strconv.FormatFloat(val, 'f', precision, 64)
}
