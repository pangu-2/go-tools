package convPg

import (
	"fmt"
	"testing"
)

//ObjToStr obj 转换为 字符
func TestObjToStr(t *testing.T) {
	obj := 64
	str := ObjToStr(obj)

	fmt.Println(str)

	var intx1 int32 = 12312
	str1 := ObjToStr(intx1)

	fmt.Println(str1)

	var float1 float64 = 121313131312312.1231234234232453242
	fmt.Println(float1)
	str2 := ObjToStr(float1)

	fmt.Println(str2)
	//

}
