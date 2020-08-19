package controller

import (
	"curve/src/model"
	"curve/src/utils"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Photo(c *gin.Context) {
	countOfPhoto, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isThumbnail := c.Query("is_thumbnail")
	datas, err := GlobalFileStorage.RandomGet(countOfPhoto, model.PhotoDirName)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isThumbnail == "true" {
		datas, err = utils.GetThumbnailDatas(datas, model.ThumbnailWidth, model.ThumbnailHeight)
		if err != nil {
			logs.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": datas,
	})
}

func UpLoadPhoto(c *gin.Context) {
	var uploadPhotoForm *model.UploadPhotoForm
	if err := c.ShouldBindJSON(&uploadPhotoForm); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := checkAndGetUid(c); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	photoNewName, err := utils.GetNewFileNameBaseOnTimeFromBase64Data(uploadPhotoForm.PhotoBase64Data)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := GlobalFileStorage.StoreBase64Data(model.PhotoDirName, photoNewName, uploadPhotoForm.PhotoBase64Data); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "图片上传成功."})
}
