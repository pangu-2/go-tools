package convPg

import "strconv"

// StrToFloat 字符转换为 float64
func StrToFloat(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

// StrToFloat32 字符转换为 float32
func StrToFloat32(str string) float32 {
	i, _ := strconv.ParseFloat(str, 32)
	return float32(i)
}
