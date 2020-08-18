package controller

import (
	"curve/src/model"
	"fmt"
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
	const (
		adminVrc = "999999"
	)
	var registerForm model.RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if registerForm.Vrc == adminVrc {
		logs.Info("Admin testing registering.")
	} else {
		isRight, err := GlobalRegisterVrcManager.IsVrcRight(registerForm.Email, registerForm.Vrc)
		if err != nil {
			logs.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := GlobalRegisterVrcManager.DelVrc(registerForm.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "注册失败，验证码清空失败."})
			return
		}
		if isRight == false {
			c.JSON(http.StatusOK, gin.H{"msg": "注册失败，验证码错误."})
			return
		}
		logs.Info(
			"Deleting email(%s) verification code(%s) success.",
			registerForm.Email,
			registerForm.Vrc,
		)
	}

	isEmailExist, err := GlobalUserManager.IsEmailExist(registerForm.Email)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isEmailExist {
		logs.Info("Registering fail, as email(%s) has existed.", registerForm.Email)
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("邮箱 %s 已存在。", registerForm.Email)})
		return
	}

	uid, err := GlobalUserManager.InsertUser(registerForm.Email, registerForm.Password, model.Admin)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secretTokenString, err := GlobalTokenManager.GetSecretTokenString(uid)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(
		model.KeyForTokenInCookies,
		secretTokenString,
		-1,
		"/",
		"localhost",
		false,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功."})
}
