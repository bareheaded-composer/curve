package main

import (
	"curve/src/controller"
	"curve/src/dao"
	"curve/src/env"
	"curve/src/handler"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"html/template"
	"regexp"
)

const (
	VrcLength = 6
	VrcPool   = "0123456789"
	confPath  = "conf/conf.json"
)

func init() {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
}

func init() {
	if err := env.Conf.Load(confPath); err != nil {
		logs.Error("Load conf(%s) failed.", confPath)
		return
	}
}

func init() {
	const (
		charPool             = "0123456789"
		vrcLength            = 6
		vrcExpiredSecond     = 60
		emailTemplateContent = `您的验证码是: {{.Vrc}} 验证码过期时间为: {{.VrcExpiredSecond}}s.`
	)
	vrcGenerator := handler.NewVrcGenerator(charPool, vrcLength)
	client := handler.NewEmailClient(
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
	vrcEmailSender := handler.NewVrcEmailSender(client, vrcGenerator, emailTemplate)
	cache := dao.NewCache(
		env.Conf.Cache.Network,
		env.Conf.Cache.Host,
	)
	controller.GlobalRegisterVrcManager = handler.NewRegisterVrcManager(vrcEmailSender, cache, vrcExpiredSecond)
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("password", password); err != nil {
			logs.Error(err)
			return
		}
		if err := v.RegisterValidation("vrc", vrc); err != nil {
			logs.Error(err)
			return
		}
	}
	r.GET("/test", controller.Test)
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/askForRegister", controller.AskForRegister)
	if err := r.Run(fmt.Sprintf(":%d", env.Conf.Http.Port)); err != nil {
		logs.Error("Running go http server failed. :|")
		return
	}
}

func password(f validator.FieldLevel) bool {
	const passwordRegexp = `^[_.a-zA-Z0-9]{8,15}$`
	reg := regexp.MustCompile(passwordRegexp)
	return reg.Match([]byte(f.Field().String()))
}

func vrc(f validator.FieldLevel) bool {
	var vrcRegexp = fmt.Sprintf(`^[%s]{%d,%d}$`, VrcPool, VrcLength, VrcLength)
	reg := regexp.MustCompile(vrcRegexp)
	return reg.Match([]byte(f.Field().String()))
}
