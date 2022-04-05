package arrayPg

import (
	"github.com/pangu-2/go-tools/tools/convPg"
)

//StringReverse slice翻转
func StringReverse(src []string) error {
	return convPg.StringReverse(src)
}
