package response

const ErrCodeParamInvalid = 20003
const ErrCodeTokenInvalid = 30001
const ErrCodeSuccess = 20001

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "param invalid",
	ErrCodeTokenInvalid: "token invalid",
}
