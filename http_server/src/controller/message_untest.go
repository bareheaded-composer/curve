package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendMessage(c *gin.Context) {
	var sendMessageForm model.SendLetterForm
	if err := c.ShouldBindJSON(&sendMessageForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	senderUID, err := getUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalMessageManager.StoreMessage(
		senderUID,
		sendMessageForm.ReceiverUID,
		sendMessageForm.Content,
	); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.Info("Storing message of sender_uid(%d) to receiver_uid(%d) success.", senderUID, sendMessageForm.ReceiverUID)
	if err := GlobalMessageManager.SendMessage(
		senderUID,
		sendMessageForm.ReceiverUID,
		sendMessageForm.Content,
	); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "消息发送成功。"})
}

func RegisterClientOfReceivingMessage(c *gin.Context) {
	uid, err := getUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalMessageManager.SetConn(uid, c); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "消息接收客户端注册成功。"})
}
