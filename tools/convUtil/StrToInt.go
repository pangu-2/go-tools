package convUtil

import "strconv"

func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
