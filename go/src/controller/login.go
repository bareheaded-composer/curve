package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Login(c *gin.Context) {
	var loginForm model.LoginForm
	if err := c.ShouldBindJSON(&loginForm);err!=nil{
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.Info(loginForm)
}
