package r

type Option func(c R)

func Of(options ...Option) R {
	opts := R{}
	opts.Code = SUCCESS_CODE
	for _, option := range options {
		option(opts)
	}
	return opts
}

// 正确
func OkDefault() R {
	return R{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE}
}

// 正确
func OkDefaultPointer() *R {
	return &R{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE}
}

// 正确
func Ok(str string) R {
	return R{Code: SUCCESS_CODE, Message: str}
}

// 正确
func OkPointer(str string) *R {
	return &R{Code: SUCCESS_CODE, Message: str}
}

// 正确
func OkData(data interface{}) R {
	return R{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE, Data: data}
}

// 正确
func OkDataPointer(data interface{}) *R {
	return &R{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE, Data: data}
}

// 正确
func OkR2(str string) (string, R) {
	return SUCCESS_CODE, R{Code: SUCCESS_CODE, Message: str}
}

// 正确
func OkR2Pointer(str string) (string, *R) {
	return SUCCESS_CODE, &R{Code: SUCCESS_CODE, Message: str}
}

// 错误
func ErrorDefault() R {
	return R{Code: ERROR_CODE, Message: ERROR_MESSAGE}
}

// 错误
func ErrorDefaultPointer() *R {
	return &R{Code: ERROR_CODE, Message: ERROR_MESSAGE}
}

// 错误
func Error(str string) R {
	return R{Code: ERROR_CODE, Message: str}
}

// 错误
func ErrorPointer(str string) *R {
	return &R{Code: ERROR_CODE, Message: str}
}

// 错误
func ErrorR2(str string) (string, R) {
	return ERROR_CODE, R{Code: ERROR_CODE, Message: str}
}

// 错误
func ErrorR2Pointer(str string) (string, *R) {
	return ERROR_CODE, &R{Code: ERROR_CODE, Message: str}
}

// 非法参数
func IllegalArgument(str string) R {
	return R{Code: ILLEGAL_ARGUMENT_CODE, Message: str}
}

// 非法参数
func IllegalArgumentPointer(str string) *R {
	return &R{Code: ILLEGAL_ARGUMENT_CODE, Message: str}
}

// 非法参数
func IllegalArgumentR2(str string) (string, R) {
	return ILLEGAL_ARGUMENT_CODE, R{Code: ILLEGAL_ARGUMENT_CODE, Message: str}
}

// 非法参数
func IllegalArgumentR2Pointer(str string) (string, *R) {
	return ILLEGAL_ARGUMENT_CODE, &R{Code: ILLEGAL_ARGUMENT_CODE, Message: str}
}

// 新数据
func Wrap(code string, str string) R {
	return R{Code: code, Message: str}
}

// string
func WrapPointer(code string, str string) *R {
	return &R{Code: code, Message: str}
}

// 新数据
func WrapR2(code string, str string) (string, R) {
	return code, R{Code: code, Message: str}
}

// 新数据
func WrapR2Pointer(code string, str string) (string, *R) {
	return code, &R{Code: code, Message: str}
}

// 新数据
func WrapData(code string, str string, data interface{}) R {
	return R{Code: code, Message: str, Data: data}
}

// 新数据
func WrapDataPointer(code string, str string, data interface{}) *R {
	return &R{Code: code, Message: str, Data: data}
}

// 新数据
func WrapDataR2(code string, str string, data interface{}) (string, R) {
	return code, R{Code: code, Message: str, Data: data}
}

// 新数据
func WrapDataR2Pointer(code string, str string, data interface{}) (string, *R) {
	return code, &R{Code: code, Message: str, Data: data}
}
