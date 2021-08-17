package successx

const defaultCode = 200
const defaultMsg = "success"

type SuccessResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewMsgSuccess(msg string, data interface{}) *SuccessResponse {
	return &SuccessResponse{Code: defaultCode, Msg: msg, Data: data}
}

func NewDefaultSuccess(data interface{}) *SuccessResponse {
	return NewMsgSuccess(defaultMsg, data)
}