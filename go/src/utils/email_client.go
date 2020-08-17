package utils

import (
	"curve/src/model"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/smtp"
)

type EmailClient struct {
	emailAddr string
	authCode  string
	smtpAddr  string
	port      int
}

func NewEmailClient(emailAddr string, authCode string, smtpAddr string, smtpPort int) *EmailClient {
	return &EmailClient{
		emailAddr: emailAddr,
		authCode:  authCode,
		smtpAddr:  smtpAddr,
		port:      smtpPort,
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
	stmpPlainAuth := smtp.PlainAuth("", e.emailAddr, e.authCode, e.smtpAddr)
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
	stmpHost := fmt.Sprintf("%s:%d", e.smtpAddr, e.port)
	if err := smtp.SendMail(stmpHost, stmpPlainAuth, e.emailAddr, emailContext.EmailAddrOfReceivers, msg); err != nil {
		logs.Error(err)
	}
	return nil
}
