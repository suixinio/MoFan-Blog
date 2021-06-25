package errmsg

const (
	success = 200
	err     = 500
	//code = 1000    用户模块错误
	err_username_used    = 1001
	err_password_wrong   = 1002
	err_user_not_exist   = 1003
	err_token_exist      = 1004
	err_token_runtime    = 1005
	err_token_wrong      = 1006
	err_token_type_wrong = 1007
	//code = 2000    文章模块错误
	//code = 3000    分类模块错误

)

var codemsg = map[int]string{
	success:              "OK",
	err:                  "FAIL",
	err_username_used:    "用户名已存在",
	err_password_wrong:   "密码错误",
	err_user_not_exist:   "用户不存在",
	err_token_exist:      "TOKEN不存在",
	err_token_runtime:    "TOKEN已过期",
	err_token_wrong:      "TOKEN不正确",
	err_token_type_wrong: "TOKEN格式错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
