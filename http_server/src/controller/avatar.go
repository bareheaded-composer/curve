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
	data, err := GlobalFileStorage.Get(model.AvatarDirName,avatarPhotoName)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Msg:  "获取头像成功.",
		Data: data,
	})
}

func UpdateAvatar(c *gin.Context) {
	var updateAvatarForm *model.UpdateAvatarForm	//  为什么指针就可以呢？而不是指针就会Panic
	if err := c.ShouldBindJSON(&updateAvatarForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	uid, err := checkAndGetUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	avatarFileName, err := utils.GetNewFileNameBaseOnTimeFromBase64Data(updateAvatarForm.AvatarBase64Data)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	if err := GlobalFileStorage.StoreBase64Data(model.AvatarDirName,avatarFileName, updateAvatarForm.AvatarBase64Data); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	if err := GlobalUserManager.UpdateAvatarFileName(uid, avatarFileName); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Msg:  "头像修改成功.",
	})
}
