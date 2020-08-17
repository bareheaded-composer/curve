package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HadSentLetter(c *gin.Context) {
	uid, err := getUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	letters, err := GlobalLetterManager.GetHadSentLetters(uid)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "获取已发送消息成功", "data": letters})
}

func HadReceivedLetter(c *gin.Context) {
	senderUIDString := c.Param("senderUID")
	senderUID, err := strconv.Atoi(senderUIDString)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	receiverUID, err := getUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	letters, err := GlobalLetterManager.GetHadReceivedLetters(senderUID, receiverUID)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "获取消息成功", "data": letters})
}

func SendLetter(c *gin.Context) {
	var sendLetterForm model.SendLetterForm
	if err := c.ShouldBindJSON(&sendLetterForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secretTokenString, err := c.Cookie(model.KeyForTokenInCookies)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	senderUID, err := GlobalTokenAnnouncer.GetUidFromSecretTokenString(secretTokenString)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalLetterManager.SendLetter(
		senderUID,
		sendLetterForm.ReceiverUID,
		sendLetterForm.Content,
	); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "消息发送成功。"})
}
