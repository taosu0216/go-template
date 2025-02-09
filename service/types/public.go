package types

type BasicResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Err  error  `json:"err"`
}
