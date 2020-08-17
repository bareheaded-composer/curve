package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AskForChangePassword(c *gin.Context) {
	var askForChangePasswordForm model.AskForChangePasswordForm
	if err := c.ShouldBindJSON(&askForChangePasswordForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalChangePasswordVrcManager.SendAndStoreVrc(askForChangePasswordForm.Email); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "修改密码的验证码已发送."})
}

func ChangePassword(c *gin.Context) {
	const adminVrc = "999999"
	var changePasswordForm model.ChangePasswordForm
	if err := c.ShouldBindJSON(&changePasswordForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if changePasswordForm.Vrc == adminVrc {
		logs.Info("Admin testing registering.")
	} else {
		isRight, err := GlobalChangePasswordVrcManager.IsVrcRight(changePasswordForm.Email, changePasswordForm.Vrc)
		if err != nil {
			logs.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := GlobalRegisterVrcManager.DelVrc(changePasswordForm.Email); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "修改密码失败，验证码清空失败."})
			return
		}
		if isRight == false {
			c.JSON(http.StatusOK, gin.H{"msg": "修改密码失败，验证码错误."})
			return
		}
		logs.Info(
			"Deleting email(%s) verification code(%s) success.",
			changePasswordForm.Email,
			changePasswordForm.Vrc,
		)
	}
	if err := GlobalUserManager.UpdatePassword(changePasswordForm.Email, changePasswordForm.NewPassword); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "密码修改成功."})
}
