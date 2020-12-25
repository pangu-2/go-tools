package R

import "net/http"

type R struct {
	Code    int         `json:"code"`    // 代码
	Message string      `json:"message"` // 信息
	Data    interface{} `json:"data"`    // 其他数据
}

//正确
func OkDefault() *R {
	return &R{Code: http.StatusOK, Message: SUCCESS_MESSAGE}
}

//正确
func Ok(str string) *R {
	return &R{Code: http.StatusOK, Message: str}
}

//正确
func OkData(data interface{}) *R {
	return &R{Code: http.StatusOK, Message: SUCCESS_MESSAGE, Data: data}
}

//正确
func OkR2(str string) (int, *R) {
	return http.StatusOK, &R{Code: http.StatusOK, Message: str}
}

//错误
func ErrorDefault() *R {
	return &R{Code: http.StatusInternalServerError, Message: ERROR_MESSAGE}
}

//错误
func Error(str string) *R {
	return &R{Code: http.StatusInternalServerError, Message: str}
}

//错误
func ErrorR2(str string) (int, *R) {
	return http.StatusInternalServerError, &R{Code: http.StatusInternalServerError, Message: str}
}

//非法参数
func IllegalArgument(str string) *R {
	return &R{Code: http.StatusUnsupportedMediaType, Message: str}
}

//非法参数
func IllegalArgumentR2(str string) (int, *R) {
	return http.StatusUnsupportedMediaType, &R{Code: http.StatusUnsupportedMediaType, Message: str}
}

//新数据
func Wrap(code int, str string) *R {
	return &R{Code: code, Message: str}
}

//新数据
func WrapR2(code int, str string) (int, *R) {
	return code, &R{Code: code, Message: str}
}

//新数据
func WrapData(code int, str string, data interface{}) *R {
	return &R{Code: code, Message: str, Data: data}
}

//新数据
func WrapDataR2(code int, str string, data interface{}) (int, *R) {
	return code, &R{Code: code, Message: str, Data: data}
}

func NewR() *R {
	r := new(R)
	r.Code = http.StatusOK
	return r
}

//是否正确
func (c *R) CodeSuccess() bool {
	return c.Code == http.StatusOK
}

//是否错误
func (c *R) CodeError() bool {
	return c.Code == http.StatusInternalServerError
}

//返回结构体
func (c *R) ResultR2() (int, *R) {
	return c.Code, c
}
