package r

type R struct {
	Code       string                 `json:"code"`                 // 代码
	Message    string                 `json:"message"`              // 信息
	Data       interface{}            `json:"data"`                 // 其他数据
	Encryption interface{}            `json:"encryption,omitempty"` // 加密数据
	Type       string                 `json:"type,omitempty"`       // 类型
	Timestamp  string                 `json:"timestamp,omitempty"`  // 时间戳
	Extend     map[string]interface{} `json:"extend,omitempty"`     // 其他数据
}

func NewR() *R {
	r := new(R)
	r.Code = SUCCESS_CODE
	return r
}

// 是否正确
func (c R) SuccessIs() bool {
	return c.Code == SUCCESS_CODE
}

// 是否错误
func (c R) ErrorIs() bool {
	return c.Code == ERROR_CODE
}

// 返回结构体
func (c R) R2String() (string, R) {
	return c.Code, c
}

// 返回结构体
func (c R) R2Bool() (bool, R) {
	return c.SuccessIs(), c
}

// 返回结构体
func (c R) Ok() R {
	c.Code = SUCCESS_CODE
	c.Message = SUCCESS_MESSAGE
	return c
}

// 返回结构体
func (c R) OkData(d interface{}) R {
	c.Ok()
	c.Data = d
	return c
}

// 返回结构体
func (c R) Error() R {
	c.Code = ERROR_CODE
	c.Message = ERROR_MESSAGE
	return c
}

// 返回结构体
func (c R) ErrorData(d interface{}) R {
	c.Error()
	c.Data = d
	return c
}

// 返回结构体
func (c R) Wrap(code string, message string, data interface{}) R {
	c.Code = code
	c.Message = message
	c.Data = data
	return c
}

// 是否正确
func (c *R) SuccessIsPointer() bool {
	return c.Code == SUCCESS_CODE
}

// 是否错误
func (c *R) ErrorIsPointer() bool {
	return c.Code == ERROR_CODE
}

// 返回结构体
func (c *R) R2PointerString() (string, *R) {
	return c.Code, c
}

// 返回结构体
func (c *R) R2PointerBool() (bool, *R) {
	return c.SuccessIsPointer(), c
}
