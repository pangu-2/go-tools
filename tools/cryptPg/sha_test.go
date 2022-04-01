package cryptPg

import (
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123456"))
	fmt.Println(Sha256("123456"))
}
