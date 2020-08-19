package controller

import (
	"curve/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, model.Response{
		Msg:"Running go http server success. :)",
	})
}
