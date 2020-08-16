package model

import "github.com/jinzhu/gorm"

const (
	tableNameOfUAI = "user_account_information"
	tableNameOfUPI = "user_personal_information"
)

type UAI struct {
	gorm.Model
	UID           int    `json:"uid"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Type          int    `json:"type"`
	RegisterTime  int    `json:"register_time"`
	LastLoginTime int    `json:"last_login_time"`
	Salt          string `json:"salt"`
}

func (*UAI) TableName() string {
	return tableNameOfUAI
}

type UPI struct {
	gorm.Model
	UID          int    `json:"uid"`
	AvatarPath   string `json:"avatar_path"`
	Username     string `json:"username"`
	Sex          int    `json:"sex"`
	ContactPhone string `json:"contact_phone"`
	ContactEmail string `json:"contact_email"`
	Birthday     int    `json:"birthday"`
}

func (*UPI) TableName() string {
	return tableNameOfUPI
}
