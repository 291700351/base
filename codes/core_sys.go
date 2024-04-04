package codes

type ResponseCode struct {
	Code int
	Msg  string
}

func NewCode(code int, msg string) *ResponseCode {
	return &ResponseCode{Code: code, Msg: msg}
}

type codeList struct {
	Success            *CodeError
	UnknownError       *CodeError
	SystemError        *CodeError
	SaveDatabaseFailed *CodeError
	DataNotExist       *CodeError
}

var SysCode = codeList{
	Success:            NewCodeError(NewCode(0, "")),
	UnknownError:       NewCodeError(NewCode(1, "系统发生未知错误，请稍候重试")),
	SystemError:        NewCodeError(NewCode(2, "系统错误，请稍候重试")),
	SaveDatabaseFailed: NewCodeError(NewCode(3, "数据保存失败")),
	DataNotExist:       NewCodeError(NewCode(4, "未找到对应数据")),
}
