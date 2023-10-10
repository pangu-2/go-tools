package strPg

import "strings"

// ReplacePrefix 替换前缀
//
//	@Description:
//	@param str
//	@param prefix
//	@return string
func ReplacePrefix(str, prefix string) string {
	if find := strings.Index(str, prefix); 0 == find {
		return strings.Replace(str, prefix, "", -1)
	}
	return str
}
