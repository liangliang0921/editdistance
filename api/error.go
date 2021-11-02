package api

const (
	ErrorUnmarshFailed = 1001
	ErrorGetRawFailed  = 1002
)

var errorMsg = map[int64]string{
	ErrorUnmarshFailed: "参数解析错误",
	ErrorGetRawFailed:  "newstring和oldstring未填",
}

func GetErrorCodeMsg(code int64) string {
	if msg, ok := errorMsg[code]; ok {
		return msg
	}
	return "未知错误"
}

func NewError(code int64) ApiResp {
	msg := GetErrorCodeMsg(code)
	return ApiResp{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
