package convPg

import "strconv"

func Float32tToStr(i float32) string {
	return strconv.FormatFloat(float64(i), 'f', 6, 64)
}

func Float64ToStr(i float64) string {
	return strconv.FormatFloat(i, 'f', 6, 64)
}
