package Rs

type Rs[T any] struct {
	Code       string                 `json:"code"`                 // 代码
	Message    string                 `json:"message"`              // 信息
	Data       T                      `json:"data"`                 // 数据
	Encryption interface{}            `json:"encryption,omitempty"` // 加密数据
	Type       string                 `json:"type,omitempty"`       // 类型
	Timestamp  string                 `json:"timestamp,omitempty"`  // 时间戳
	Extend     map[string]interface{} `json:"extend,omitempty"`     // 其他数据
}

type Option[T any] func(c Rs[T])
type OptionPointer[T any] func(c *Rs[T])

// 正确
func OkDefault[T any]() Rs[T] {
	return Rs[T]{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE}
}

// 正确
func OkDefaultPointer[T any]() *Rs[T] {
	return &Rs[T]{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE}
}

// 正确
func Ok[T any](str string) Rs[T] {
	return Rs[T]{Code: SUCCESS_CODE, Message: str}
}

// 正确
func OkPointer[T any](str string) *Rs[T] {
	return &Rs[T]{Code: SUCCESS_CODE, Message: str}
}

// 正确
func OkData[T any](data T) Rs[T] {
	return Rs[T]{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE, Data: data}
}

// 正确
func OkDataPointer[T any](data T) *Rs[T] {
	return &Rs[T]{Code: SUCCESS_CODE, Message: SUCCESS_MESSAGE, Data: data}
}

// 正确
func OkR2[T any](str string) (string, Rs[T]) {
	return SUCCESS_CODE, Rs[T]{Code: SUCCESS_CODE, Message: str}
}

// 正确
func OkR2Pointer[T any](str string) (string, *Rs[T]) {
	return SUCCESS_CODE, &Rs[T]{Code: SUCCESS_CODE, Message: str}
}

// 错误
func ErrorDefault[T any]() Rs[T] {
	return Rs[T]{Code: ERROR_CODE, Message: ERROR_MESSAGE}
}

// 错误
func ErrorDefaultPointer[T any]() *Rs[T] {
	return &Rs[T]{Code: ERROR_CODE, Message: ERROR_MESSAGE}
}

// 错误
func Error[T any](str string) Rs[T] {
	return Rs[T]{Code: ERROR_CODE, Message: str}
}

// 错误
func ErrorPointer[T any](str string) *Rs[T] {
	return &Rs[T]{Code: ERROR_CODE, Message: str}
}

// 错误
func ErrorR2[T any](str string) (string, Rs[T]) {
	return ERROR_CODE, Rs[T]{Code: ERROR_CODE, Message: str}
}

// 错误
func ErrorR2Pointer[T any](str string) (string, *Rs[T]) {
	return ERROR_CODE, &Rs[T]{Code: ERROR_CODE, Message: str}
}

// 非法参数
func IllegalArgument[T any](str string) Rs[T] {
	return Rs[T]{Code: ILLEGAL_ARGUMENT_CODE_, Message: str}
}

// 非法参数
func IllegalArgumentPointer[T any](str string) *Rs[T] {
	return &Rs[T]{Code: ILLEGAL_ARGUMENT_CODE_, Message: str}
}

// 非法参数
func IllegalArgumentR2[T any](str string) (string, Rs[T]) {
	return ILLEGAL_ARGUMENT_CODE_, Rs[T]{Code: ILLEGAL_ARGUMENT_CODE_, Message: str}
}

// 非法参数
func IllegalArgumentR2Pointer[T any](str string) (string, *Rs[T]) {
	return ILLEGAL_ARGUMENT_CODE_, &Rs[T]{Code: ILLEGAL_ARGUMENT_CODE_, Message: str}
}

// 新数据
func Wrap[T any](code string, str string) Rs[T] {
	return Rs[T]{Code: code, Message: str}
}

// 新数据
func WrapPointer[T any](code string, str string) *Rs[T] {
	return &Rs[T]{Code: code, Message: str}
}

// 新数据
func WrapR2[T any](code string, str string) (string, Rs[T]) {
	return code, Rs[T]{Code: code, Message: str}
}

// 新数据
func WrapR2Pointer[T any](code string, str string) (string, *Rs[T]) {
	return code, &Rs[T]{Code: code, Message: str}
}

// 新数据
func WrapData[T any](code string, str string, data interface{}) Rs[T] {
	return Rs[T]{Code: code, Message: str, Data: data}
}

// 新数据
func WrapDataPointer[T any](code string, str string, data interface{}) *Rs[T] {
	return &Rs[T]{Code: code, Message: str, Data: data}
}

// 新数据
func WrapDataR2[T any](code string, str string, data interface{}) (string, Rs[T]) {
	return code, Rs[T]{Code: code, Message: str, Data: data}
}

// 新数据
func WrapDataR2Pointer[T any](code string, str string, data interface{}) (string, *Rs[T]) {
	return code, &Rs[T]{Code: code, Message: str, Data: data}
}

func OfPointer[T any](options ...OptionPointer[T]) *Rs[T] {
	opts := new(Rs[T])
	opts.Code = SUCCESS_CODE
	for _, option := range options {
		option(opts)
	}
	return opts
}
func Of[T any](options ...Option[T]) Rs[T] {
	opts := Rs[T]{}
	opts.Code = SUCCESS_CODE
	for _, option := range options {
		option(opts)
	}
	return opts
}

// 是否正确
func (c Rs[T]) CodeSuccess() bool {
	return c.Code == SUCCESS_CODE
}

// 是否错误
func (c Rs[T]) CodeError() bool {
	return c.Code == ERROR_CODE
}

// 返回结构体
func (c Rs[T]) ResultR2() (string, Rs[T]) {
	return c.Code, c
}

// 是否正确
func (c *Rs[T]) CodeSuccessPointer() bool {
	return c.Code == SUCCESS_CODE
}

// 是否错误
func (c *Rs[T]) CodeErrorPointer() bool {
	return c.Code == ERROR_CODE
}

// 返回结构体
func (c *Rs[T]) ResultR2Pointer() (string, *Rs[T]) {
	return c.Code, c
}
