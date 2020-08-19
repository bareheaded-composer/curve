package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	tableNameOfUAI     = "user_account_information"
	tableNameOfUPI     = "user_personal_information"
	tableNameOfLetter  = "letter"
	tableNameOfMessage = "message"
)

type UAI struct {
	gorm.Model
	Email             string    `json:"email"`
	HashSaltyPassword string    `json:"hash_salty_password"`
	Type              UserType  `json:"type"`
	LastLoginTime     time.Time `json:"last_login_time"`
	Salt              string    `json:"salt"`
}

func (*UAI) TableName() string {
	return tableNameOfUAI
}

type UPI struct {
	gorm.Model
	AvatarPath   string    `json:"avatar_path"`
	Username     string    `json:"username"`
	Sex          int       `json:"sex"`
	ContactPhone string    `json:"contact_phone"`
	ContactEmail string    `json:"contact_email"`
	Birthday     time.Time `json:"birthday"`
}

func (*UPI) TableName() string {
	return tableNameOfUPI
}

type Letter struct {
	gorm.Model
	SenderUID   int
	ReceiverUID int
	Content     string
}

func (*Letter) TableName() string {
	return tableNameOfLetter
}

type Message struct {
	gorm.Model
	SenderUID   int
	ReceiverUID int
	Content     string
}

func (*Message) TableName() string {
	return tableNameOfMessage
}

