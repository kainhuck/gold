package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_EXIST_USERNAME:           "用户名已经存在",
	ERROR_DIFF_PASSWORD:            "两次密码不一致",
	ERROR_NOT_EXIST_USERNAME:       "用户名不存在",
	ERROR_INCORRECT_PASSWORD:       "密码错误",
	ERROR_EDIT_PASSWORD:            "密码修改失败",
	ERROR_NOT_EXIST_USERID:         "用户ID不存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
