package convPg

import (
	"fmt"
	"testing"
)

func TestNumberToStr(t *testing.T) {
	str := NumberToStr(60)
	fmt.Println(str)
	fmt.Println(str)
	str2 := NumberToStrDefault(0, "01")
	fmt.Println(str2)
}
