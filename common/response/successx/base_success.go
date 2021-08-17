package successx

const DefaultCode = 200
const DefaultMsg = "success"

type SuccessResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewMsgSuccess(msg string, data interface{}) *SuccessResponse {
	return &SuccessResponse{Code: DefaultCode, Msg: msg, Data: data}
}

func NewDefaultSuccess(data interface{}) *SuccessResponse {
	return NewMsgSuccess(DefaultMsg, data)
}