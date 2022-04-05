package numberPg

import (
	"strconv"

	"github.com/pangu-2/go-tools/tools/convPg"
)

//IsInt0 检测 是否为0
func IsInt0(val int) bool {
	return 0 == val
}

//此处只记录 转换方式，并不使用
//
//数字变成字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

//数字变成字符串
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ObjToInt 转成数字
func ObjToInt(val interface{}) int {
	return convPg.ObjToInt(val)
}

// ObjToInt32 转成数字
func ObjToInt32(val interface{}) int32 {
	return convPg.ObjToInt32(val)
}

// ObjToInt64 转成数字
func ObjToInt64(val interface{}) int64 {
	return convPg.ObjToInt64(val)
}

// ObjToFloat32 转成数字
func ObjToFloat32(val interface{}) float32 {
	return convPg.ObjToFloat32(val)
}

// ObjToFloat32 转成数字
func ObjToFloat64(val interface{}) float64 {
	return convPg.ObjToFloat64(val)
}

//字符串转换成  int 没有错误返回
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//字符串转换成  int64 没有错误返回
func StrToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

//字符串转换成  float64 没有错误返回
func StrToFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

//字符串转换成  float64 没有错误返回
func StrToFloat64FormInterface(str interface{}) float64 {
	i, _ := strconv.ParseFloat(str.(string), 64)
	return i
}

//字符串转换成  int 没有错误返回
func StrToIntFormInterface(str interface{}) int {
	i, _ := strconv.Atoi(str.(string))
	return i
}
