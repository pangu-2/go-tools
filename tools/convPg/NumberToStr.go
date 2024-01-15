package convPg

import (
	"reflect"
	"strconv"

	"github.com/pangu-2/go-tools/generics"
)

// NumberToStr 数值转换为字符
func NumberToStr[V generics.Signed](val V) string {
	return NumberToStrDefault(val, "")
}

// NumberToStrPrec
//
//	@Description:数值转换为字符
//	@param val
//	@param precision 精度
//	@return string
func NumberToStrPrec[V generics.Signed](val V, precision int) string {
	return NumberToStrDefaultPrec[V](val, "", precision)
}

// NumberToStrDefault 数值转换为字符
func NumberToStrDefault[V generics.Number](val V, def string) string {
	return NumberToStrDefaultPrec(val, def, -1)
}

// NumberToStrDefaultPrec
//
//	@Description: 数值转换为字符
//	@param val
//	@param def 默认值
//	@param precision 精度
//	@return string
func NumberToStrDefaultPrec[V generics.Number](val V, def string, precision int) string {
	rv := reflect.TypeOf(val)
	//fmt.Print(of.String())
	//fmt.Print(of.Key())
	// 显示反射类型对象的名称和种类
	//fmt.Println(of.Name(), of.Kind())
	str := ""
	switch rv.Kind() {
	case reflect.Int:
		str = strconv.Itoa(int(val))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = strconv.FormatInt(int64(val), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		str = strconv.FormatUint(uint64(val), 10)
	case reflect.Float32:
		str = strconv.FormatFloat(float64(val), 'f', precision, 64)
	case reflect.Float64:
		str = strconv.FormatFloat(float64(val), 'f', precision, 64)
	}
	//是否为空
	if len(str) > 0 {
		return str
	}
	return def
}
