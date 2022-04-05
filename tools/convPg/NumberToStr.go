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

// NumberToStrDefault 数值转换为字符
func NumberToStrDefault[V generics.Number](val V, def string) string {
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
		str = strconv.FormatFloat(float64(val), 'f', 6, 64)
	case reflect.Float64:
		str = strconv.FormatFloat(float64(val), 'f', 6, 64)
	}
	//是否为空
	if len(str) > 0 {
		return str
	}
	return def
}
