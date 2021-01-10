package structUtil

import "github.com/gokits/gotools"

//demo : https://github.com/gokits/gotools
//复制
//src 原
func StructCopy(src interface{}, dst interface{}) {
	gotools.StructCopy(dst, src)
}
