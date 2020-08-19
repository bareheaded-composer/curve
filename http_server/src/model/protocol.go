package model

type Response struct {
	Code int         // 业务自定义状态码
	Msg  string      // 返回消息，包括错误消息，成功消息
	data interface{}
}
