package model

type Response struct {
	Err string		 // 错误消息
	Msg  string      // 返回消息，成功消息
	data interface{}
}
