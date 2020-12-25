package strUtil

import (
	"fmt"
	"strings"
)


//https://studygolang.com/articles/4287
//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if length < 1 {
		return str
	}
	if start < 0 || start > length {
		fmt.Println("Substr error: start is wrong")
		return str
		//panic("start is wrong")
	}

	if end < 0 {
		fmt.Println("Substr error: end is wrong")
		return str
		//panic("end is wrong")
	}
	if end > length {
		end = length
	}

	return string(rs[start:end])
}

//替换
func Replace(str string, old, new string) string {
	return strings.Replace(str, old, new, -1)
}

//字符是否存在
func Exist(str, find string) bool {
	return strings.LastIndex(str, find) != -1
}

//字符是否存在
func LastIndex(str, find string) bool {
	return strings.LastIndex(str, find) != -1
}

//字符串比较 是否相等  ，这里加入，防止以后忘记
func EqualFold(str1, str2 string) bool {
	return strings.EqualFold(str1, str2)
}