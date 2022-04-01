package convPg

import (
	"github.com/pangu-2/go-tools/generics"
	"reflect"
	"strconv"
)

// NumberToStr 数值转换为字符
func NumberToStr[V generics.Signed](i V) string {
	return NumberToStrDefault(i, "")
}

// NumberToStrDefault 数值转换为字符
func NumberToStrDefault[V generics.Number](i V, def string) string {
	of := reflect.TypeOf(i)
	//fmt.Print(of.String())
	//fmt.Print(of.Key())
	// 显示反射类型对象的名称和种类
	//fmt.Println(of.Name(), of.Kind())
	str := ""
	switch of.Kind() {
	case reflect.Int:
		str = strconv.Itoa(int(i))
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		str = strconv.FormatInt(int64(i), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		str = strconv.FormatUint(uint64(i), 10)
	case reflect.Float32:
		str = strconv.FormatFloat(float64(i), 'f', 6, 64)
	case reflect.Float64:
		str = strconv.FormatFloat(float64(i), 'f', 6, 64)
	}
	//是否为空
	if len(str) > 0 {
		return str
	}
	return def
}
