package handler

import (
	"curve/src/model"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/smtp"
)


type EmailClient struct {
	emailAddr string
	authCode  string
	stmpAddr  string
	port      int
}

func NewEmailClient(emailAddr string, authCode string, stmpAddr string, stmpPort int) *EmailClient {
	return &EmailClient{
		emailAddr: emailAddr,
		authCode:  authCode,
		stmpAddr:  stmpAddr,
		port:      stmpPort,
	}
}

func (e *EmailClient) SendEmail(emailAddrOfReceiver string, subject string, htmlContent string) error {
	if err := e.send(&model.EmailContext{
		EmailAddrOfSender:    e.emailAddr,
		EmailAddrOfReceivers: []string{emailAddrOfReceiver},
		Subject:              subject,
		Body:                 htmlContent,
		Type:                 model.HtmlType,
	}); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (e *EmailClient) send(emailContext *model.EmailContext) error {
	stmpPlainAuth := smtp.PlainAuth("", e.emailAddr, e.authCode, e.stmpAddr)
	contentType, err := emailContext.GetSendingContentType()
	if err != nil {
		logs.Error(err)
		return err
	}
	msg := []byte(
		fmt.Sprintf("From: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s",
			emailContext.EmailAddrOfSender,
			emailContext.Subject,
			contentType,
			emailContext.Body,
		),
	)
	stmpHost := fmt.Sprintf("%s:%d", e.stmpAddr, e.port)
	if err := smtp.SendMail(stmpHost, stmpPlainAuth, e.emailAddr, emailContext.EmailAddrOfReceivers, msg); err != nil {
		logs.Error(err)
	}
	return nil
}
