package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AskForRegister(c *gin.Context) {
	var askForRegisterForm model.AskForRegisterForm
	if err := c.ShouldBindJSON(&askForRegisterForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalRegisterVrcManager.SendAndStoreVrc(askForRegisterForm.Email); err != nil {
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
	isRight, err := GlobalRegisterVrcManager.IsVrcRight(registerForm.Email, registerForm.Vrc)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalRegisterVrcManager.DelVrcOfRegisterEmail(registerForm.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "注册失败，验证码清空失败."})
		return
	}
	if isRight == false {
		c.JSON(http.StatusOK, gin.H{"msg": "注册失败，验证码错误."})
		return
	}
	logs.Info(
		"Registering(%s) success. Deleting verification code(%s) success.",
		registerForm.Email,
		registerForm.Vrc,
	)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功."})
}
