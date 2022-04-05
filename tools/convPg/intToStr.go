package convPg

import (
	"strconv"
)

func IntToStr(val int) string {
	return strconv.Itoa(val)
}

func Int8ToStr(val int8) string {
	return strconv.FormatInt(int64(val), 10)
}

func Int16ToStr(val int16) string {
	return strconv.FormatInt(int64(val), 10)
}

func Int32ToStr(val int32) string {
	return strconv.FormatInt(int64(val), 10)
}

func Int64ToStr(val int64) string {
	return strconv.FormatInt(val, 10)
}
