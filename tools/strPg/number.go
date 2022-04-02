package strPg

import "github.com/pangu-2/go-tools/tools/numberPg"

//字符串转换成  int 没有错误返回
func ToInt(str string) int {
	return numberPg.StrToInt(str)
}

//字符串转换成  int64 没有错误返回
func ToInt64(str string) int64 {
	return numberPg.StrToInt64(str)
}

//字符串转换成  float64 没有错误返回
func ToFloat64(str string) float64 {
	return numberPg.StrToFloat64(str)
}

//字符串转换成  float64 没有错误返回
func ToFloat64FormInterface(str interface{}) float64 {
	return numberPg.StrToFloat64FormInterface(str)
}

//字符串转换成  int 没有错误返回
func ToIntFormInterface(str interface{}) int {
	return numberPg.StrToIntFormInterface(str)
}
