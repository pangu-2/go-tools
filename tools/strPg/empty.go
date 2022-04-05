package strPg

import "strings"

// IsEmpty 验证是否为空
func IsEmpty(val string) bool {
	return "" == val
}

// IsNotEmpty 验证 不为空
func IsNotEmpty(val string) bool {
	return "" != val
}

// IsBlank 验证是否为空,去除左右空格
func IsBlank(val string) bool {
	return "" == strings.TrimSpace(val)
}

// IsNotBlank 验证 不为空,去除左右空格
func IsNotBlank(val string) bool {
	return "" != strings.TrimSpace(val)
}
