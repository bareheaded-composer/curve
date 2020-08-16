package handler

import (
	"curve/src/env"
	"github.com/astaxie/beego/logs"
	"html/template"
	"testing"
)

func TestVrcEmailSender(t *testing.T) {
	const charPool = "0123456789"
	const vrcLength = 6
	const emailTemplateContent = `您的验证码是: {{.Vrc}} 验证码过期时间为: {{.VrcExpiredSecond}}s.`
	vrcGenerator := NewVrcGenerator(charPool, vrcLength)
	client := NewEmailClient(
		env.Conf.EmailClient.EmailAddr,
		env.Conf.EmailClient.AuthCode,
		env.Conf.EmailClient.SmtpAddr,
		env.Conf.EmailClient.SmtpPort,
	)
	emailTemplate := template.New("")
	if _, err := emailTemplate.Parse(emailTemplateContent); err != nil {
		logs.Error(err)
		return
	}
	vrcEmailSender := NewVrcEmailSender(client, vrcGenerator, emailTemplate)
	if _, err := vrcEmailSender.SendVrcEmail("测试邮件", "417165709@qq.com", 120); err != nil {
		logs.Error(err)
		return
	}
}
