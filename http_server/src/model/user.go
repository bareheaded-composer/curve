package model

type UserType = int

const (
	Admin = iota + 1
	Vip
	GeneralUser
)

type SexType int

const (
	Man = iota + 1
	Woman
	Secret
)

const (
	DefaultAvatarPath   = ""
	DefaultUsername     = "hello_word"
	DefaultSex          = Man
	DefaultContactPhone = "18946988888"
	DefaultContactEmail = "123456789@qq.com"
)
