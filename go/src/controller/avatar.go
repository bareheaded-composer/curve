package controller

import (
	"curve/src/model"
	"curve/src/utils"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Avatar(c *gin.Context) {
	avatarPhotoName := c.Param("name")
	logs.Debug(avatarPhotoName)
	data, err := GlobalFileStorage.Get(avatarPhotoName)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func UpdateAvatar(c *gin.Context) {
	var updateAvatarForm *model.UpdateAvatarForm
	if err := c.ShouldBindJSON(&updateAvatarForm); err != nil {
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
	uid, err := GlobalTokenAnnouncer.GetUidFromSecretTokenString(secretTokenString)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	avatarFileName, err := utils.GetNewFileNameBaseOnTimeFromBase64Data(updateAvatarForm.AvatarBase64Data)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalFileStorage.StoreBase64Data(avatarFileName, updateAvatarForm.AvatarBase64Data); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalUserManager.UpdateAvatarFileName(uid, avatarFileName); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "头像修改成功."})
}
