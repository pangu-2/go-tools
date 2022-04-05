package convPg

import "strconv"

func Float32tToStr(val float32) string {
	return strconv.FormatFloat(float64(val), 'f', 6, 64)
}

func Float64ToStr(val float64) string {
	return strconv.FormatFloat(val, 'f', 6, 64)
}
