package model

type Response struct {
	Err  string      `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
