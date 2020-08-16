package main

import (
	"curve/src/controller"
	"curve/src/env"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const confPath = "conf/conf.json"

const (
	VrcLength = 6
	VrcPool   = "0123456789"
)

func init() {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
}

func main() {
	if err := env.Conf.Load(confPath); err != nil {
		logs.Error("Load conf(%s) failed.", confPath)
		return
	}
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
