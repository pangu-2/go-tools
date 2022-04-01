package convPg

import "strconv"

// StrToInt 字符转换为 int
func StrToInt(str string) int {
	//一个 int 的大小是基于操作系统特定实现的，它可能是 32 位，也可能是 64 位
	i, _ := strconv.Atoi(str)
	return i
}

// StrToInt8 字符转换为 int8
func StrToInt8(str string) int8 {
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0
	}
	return int8(i)
}

// StrToInt16 字符转换为 int16
func StrToInt16(str string) int16 {
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0
	}
	return int16(i)
}

// StrToInt32 字符转换为 int32
func StrToInt32(str string) int32 {
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(i)
}

// StrToInt64 字符转换为 int64
func StrToInt64(str string) int64 {
	//参数1 数字的字符串形式
	//参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
	//参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
