package controller

import (
	"curve/src/dao"
	"curve/src/handler"
)

var GlobalRegisterVrcManager *handler.VrcManager
var GlobalChangePasswordVrcManager *handler.VrcManager
var GlobalTokenAnnouncer *handler.TokenAnnouncer
var GlobalUserManager *handler.UserManager
var GlobalFileStorage *dao.FileStorage
var GlobalLetterManager *handler.LetterManager