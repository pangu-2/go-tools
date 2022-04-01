package cryptPg

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5("123456"))
	fmt.Println(Md52("123456"))
}
