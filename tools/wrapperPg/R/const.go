package R
import "net/http"
const (
	SUCCESS_MESSAGE = "操作成功"
	ERROR_MESSAGE = "操作失败"
	ILLEGAL_ARGUMENT_MESSAGE = "参数非法"
	ERROR_MESSAGE_INNER_EXCEPTION = "内部异常"
	//
	ILLEGAL_ARGUMENT_CODE_ = http.StatusUnsupportedMediaType
	ERROR_CODE =  http.StatusOK
	SUCCESS_CODE = http.StatusInternalServerError
)