package controller

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Attend(c *gin.Context) {
	var attendForm *model.AttendForm
	if err := c.ShouldBindJSON(&attendForm); err != nil {
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
	if err := GlobalAttentionManager.StoreAttentionIfNotExist(uid, attendForm.AttendeeUID); err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "关注成功."})
}

func GetAttentionsOfAttendee(c *gin.Context) {
	uid, err := checkAndGetUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	attentions, err := GlobalAttentionManager.GetAttentionsBaseOnAttendee(uid)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Msg:  "获取关注粉丝信息成功.",
		Data: attentions,
	})
}

func GetAttentionsOfAttender(c *gin.Context) {
	uid, err := checkAndGetUid(c)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	attentions, err := GlobalAttentionManager.GetAttentionsBaseOnAttender(uid)
	if err != nil {
		logs.Error(err)
		c.JSON(http.StatusBadRequest, model.Response{
			Err: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Msg:  "获取关注偶像信息成功.",
		Data: attentions,
	})
}
