package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	tableNameOfUserInformation = "user_information"
	tableNameOfLetter          = "letter"
	tableNameOfMessage         = "message"
	tableNameOfAttention       = "attention"
)

type UserInformation struct {
	gorm.Model
	Email             string    `json:"email"`
	HashSaltyPassword string    `json:"hash_salty_password"`
	Type              UserType  `json:"type"`
	LastLoginTime     time.Time `json:"last_login_time"`
	Salt              string    `json:"salt"`
	AvatarPath        string    `json:"avatar_path"`
	Username          string    `json:"username"`
	Sex               int       `json:"sex"`
	ContactPhone      string    `json:"contact_phone"`
	ContactEmail      string    `json:"contact_email"`
	Birthday          time.Time `json:"birthday"`
}

func (*UserInformation) TableName() string {
	return tableNameOfUserInformation
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

type Attention struct {
	gorm.Model
	AttenderUID int
	AttendeeUID int
}

func (*Attention) TableName() string {
	return tableNameOfAttention
}
