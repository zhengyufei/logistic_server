package server

type ErrorCode int

const (
	ErrorCodeNull ErrorCode = iota
	ErrorCodeInternalError
	ErrorCodeWrongParam
	ErrorCodeBadRequest
	ErrorCodeNoSign
)

var codeString = map[ErrorCode]string{
	ErrorCodeNull:          "",
	ErrorCodeInternalError: "系统错误",
	ErrorCodeWrongParam:    "参数错误",
	ErrorCodeBadRequest:    "请求错误",
	ErrorCodeNoSign:        "未认证",
}

func (ec ErrorCode) String() string {
	str, ok := codeString[ec]
	if !ok {
		return "未知错误"
	}
	return str
}
