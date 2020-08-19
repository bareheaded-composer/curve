package model

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)

const (
	RegisterEmailSubject       = "注册邮件"
	ChangePasswordEmailSubject = "修改密码邮件"
	RegisterKeyPrefix          = "register"
	ChangePasswordKeyPrefix    = "changePassword"
)

type TypeOfEmailContext string

const (
	HtmlType  TypeOfEmailContext = "html"
	PlainType TypeOfEmailContext = "plain"
)

type EmailContext struct {
	EmailAddrOfSender    string
	EmailAddrOfReceivers []string
	Subject              string
	Body                 string
	Type                 TypeOfEmailContext
}

func (c *EmailContext) GetSendingContentType() (string, error) {
	var contentType string
	switch c.Type {
	case HtmlType:
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", c.Type)
	case PlainType:
		contentType = fmt.Sprintf("text/%s; charset=UTF-8", c.Type)
	default:
		err := fmt.Errorf(fmt.Sprintf("the type(%s) of the email can't be distinguished", c.Type))
		logs.Error(err)
		return "", err
	}
	return contentType, nil
}
