package convPg

import (
	"strconv"
)

func IntToStr(i int) string {
	return strconv.Itoa(i)
}

func Int8ToStr(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int16ToStr(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int32ToStr(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
