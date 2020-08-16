package handler

import (
	"bytes"
	"github.com/astaxie/beego/logs"
	"html/template"
)

type VrcEmailSender struct {
	emailClient   *EmailClient
	vrcGenerator  *VrcGenerator
	emailTemplate *template.Template
}

func NewVrcEmailSender(client *EmailClient, generator *VrcGenerator, emailTemplate *template.Template) *VrcEmailSender {
	return &VrcEmailSender{
		emailClient:   client,
		vrcGenerator:  generator,
		emailTemplate: emailTemplate,
	}
}

func (v *VrcEmailSender) SendVrcEmail(subject string, emailAddrOfReceiver string, vrcExpiredSecond int) (string, error) {
	var htmlContentBuffer bytes.Buffer
	vrc := v.vrcGenerator.GetVrc()
	if err := v.emailTemplate.Execute(&htmlContentBuffer, struct {
		Vrc            string
		VrcExpiredSecond int
	}{
		Vrc:            vrc,
		VrcExpiredSecond:vrcExpiredSecond,
	}); err != nil {
		logs.Error(err)
		return "", err
	}
	if err := v.emailClient.SendEmail(emailAddrOfReceiver, subject, htmlContentBuffer.String()); err != nil {
		logs.Error(err)
		return "", err
	}
	return vrc, nil
}
