package codes

func err(code int, msg string) *CodeError {
	return NewCodeError(NewCode(code, msg))
}

type userErrList struct {
	ParamCanNotNull *CodeError
	LoginFail       *CodeError
	UserExists      *CodeError
	UserDisable     *CodeError
	UserDeleted     *CodeError
}

var ErrUsers = &userErrList{
	ParamCanNotNull: err(1000, "用户名或密码不能为空"),
	LoginFail:       err(1001, "用户名或密码错误"),
	UserExists:      err(1002, "用户已存在"),
	UserDisable:     err(1003, "用户已被禁用"),
	UserDeleted:     err(1003, "用户已被删除"),
}
