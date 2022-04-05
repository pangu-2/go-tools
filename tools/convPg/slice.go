package convPg

import "github.com/pangu-2/go-tools/tools"

//StringReverse slice翻转
func StringReverse(src []string) error {
	if src == nil {
		//panic(fmt.Errorf("the src can't be empty!"))
		return tools.NewError("数组不能为空")
	}
	count := len(src)
	if count < 1 {
		return tools.NewError("数组不能为空")
	}
	mid := count / 2
	for i := 0; i < mid; i++ {
		tmp := src[i]
		src[i] = src[count-1]
		src[count-1] = tmp
		count--
	}
	return nil
}
