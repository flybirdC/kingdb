package kingdb

//msg
const (

	SuccessMsg = "Success!"
	RpcError = "the rpc error 500 !"
)

//code
const (

	StatusSuccess = 200
	StatusError = 500
)

type StatusCode struct {

	CodeState int64 `json:"code_state"`
	CodeMSG string `json:"code_msg"`
}




