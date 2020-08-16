package controller

import (
	"curve/src/dao"
	"curve/src/handler"
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

var globalRegisterVrcManager *handler.RegisterVrcManager

func init() {
	const (
		QQStmpAddr           = "smtp.qq.com"
		QQStmpPort           = 587
		WyStmpAddr           = "smtp.163.com"
		WyStmpPort           = 25
		charPool             = "0123456789"
		vrcLength            = 6
		emailAddr            = ""
		authCode             = ""
		emailTemplateContent = `您的验证码是: {{.Vrc}} 验证码过期时间为: {{.VrcExpiredSecond}}s.`
	)

	vrcGenerator := handler.NewVrcGenerator(charPool, vrcLength)
	client := handler.NewEmailClient(emailAddr, authCode, QQStmpAddr, QQStmpPort)
	emailTemplate := template.New("")
	if _, err := emailTemplate.Parse(emailTemplateContent); err != nil {
		logs.Error(err)
		return
	}
	vrcEmailSender := handler.NewVrcEmailSender(client, vrcGenerator, emailTemplate)

	const (
		cacheHost        = "127.0.0.1"
		cachePort        = 6379
		network          = "tcp"
		vrcExpiredSecond = 60
	)
	cache := dao.NewCache(network, cacheHost, cachePort)
	globalRegisterVrcManager = handler.NewRegisterVrcManager(vrcEmailSender, cache, vrcExpiredSecond)

}

func AskForRegister(c *gin.Context) {
	var askForRegisterForm model.AskForRegisterForm
	if err := c.ShouldBindJSON(&askForRegisterForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := globalRegisterVrcManager.SendAndStoreVrc(askForRegisterForm.Email); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "验证码已发送."})
}

func Register(c *gin.Context) {
	var registerForm model.RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.Info(registerForm)
}
