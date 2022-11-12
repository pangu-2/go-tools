package rs

type Rs[T any | interface{}] struct {
	Code       string                 `json:"code"`                 // 代码
	Message    string                 `json:"message"`              // 信息
	Data       T                      `json:"data"`                 // 数据
	Encryption interface{}            `json:"encryption,omitempty"` // 加密数据
	Type       string                 `json:"type,omitempty"`       // 类型
	Timestamp  string                 `json:"timestamp,omitempty"`  // 时间戳
	Extend     map[string]interface{} `json:"extend,omitempty"`     // 其他数据
}

// 是否正确
func (c Rs[T]) SuccessIs() bool {
	return c.Code == SUCCESS_CODE
}

// 是否错误
func (c Rs[T]) ErrorIs() bool {
	return c.Code == ERROR_CODE
}

// 返回结构体
func (c Rs[T]) ResultRs2() (string, Rs[T]) {
	return c.Code, c
}

// 返回结构体
func (c Rs[T]) Ok() Rs[T] {
	c.Code = SUCCESS_CODE
	c.Message = SUCCESS_MESSAGE
	return c
}

// 返回结构体
func (c Rs[T]) OkData(data T) Rs[T] {
	c.Ok()
	c.Data = data
	return c
}

// 返回结构体
func (c Rs[T]) Error() Rs[T] {
	c.Code = ERROR_CODE
	c.Message = ERROR_MESSAGE
	return c
}

// 返回结构体
func (c Rs[T]) ErrorData(data T) Rs[T] {
	c.Error()
	c.Data = data
	return c
}

// 返回结构体
func (c Rs[T]) Wrap(code string, message string, data T) Rs[T] {
	c.Code = code
	c.Message = message
	c.Data = data
	return c
}

// 是否正确
func (c *Rs[T]) SuccessPointer() bool {
	return c.Code == SUCCESS_CODE
}

// 是否错误
func (c *Rs[T]) ErrorPointer() bool {
	return c.Code == ERROR_CODE
}

// 返回结构体
func (c *Rs[T]) ResultRsPointer() (string, *Rs[T]) {
	return c.Code, c
}
