package numberPg

import (
	"github.com/pangu-2/go-tools/tools/convPg"
	"regexp"
	"strconv"
)

const PatternsInt string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"

// IsInt
//
//	@Description: 是否整数
//	@param str
//	@return bool
func IsInt(str string) bool {
	if len(str) == 0 {
		return false
	}
	return regexp.MustCompile(PatternsInt).MatchString(str)
}

// IsInt0 检测 是否为0
func IsInt0(val int) bool {
	return 0 == val
}

//	此处只记录 转换方式，并不使用
//
// IntToString 数字变成字符串
func IntToString(i int) string {
	return strconv.Itoa(i)
}

// Int8ToString 数字变成字符串
func Int8ToString(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

// Int32ToString 数字变成字符串
func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// Int64ToString 数字变成字符串
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// UInt64ToString 数字变成字符串
func UInt64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
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

// ObjToFloat64 转成数字
func ObjToFloat64(val interface{}) float64 {
	return convPg.ObjToFloat64(val)
}

// StrToInt 字符串转换成  int 没有错误返回
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// StrToInt64 字符串转换成  int64 没有错误返回
func StrToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

// StrToUInt64 字符串转换成  int64 没有错误返回
func StrToUInt64(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}

// StrToFloat64 字符串转换成  float64 没有错误返回
func StrToFloat64(str string) float64 {
	i, _ := strconv.ParseFloat(str, 64)
	return i
}

// StrToFloat64FormInterface 字符串转换成  float64 没有错误返回
func StrToFloat64FormInterface(str interface{}) float64 {
	i, _ := strconv.ParseFloat(str.(string), 64)
	return i
}

// StrToIntFormInterface 字符串转换成  int 没有错误返回
func StrToIntFormInterface(str interface{}) int {
	i, _ := strconv.Atoi(str.(string))
	return i
}

// IsIntType 类型
func IsIntType(value interface{}) bool {
	switch value.(type) {
	case int, *int, int8, *int8, int16, *int16, int32, *int32, int64, *int64:
		return true
	}
	return false
}

// IsUintType 类型
func IsUintType(value interface{}) bool {
	switch value.(type) {
	case uint, *uint, uint8, *uint8, uint16, *uint16, uint32, *uint32, uint64, *uint64:
		return true
	}
	return false
}
