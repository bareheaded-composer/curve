package controller

import (
	"curve/src/dao"
	"curve/src/handler"
)

var GlobalRegisterVrcManager *handler.VrcManager
var GlobalChangePasswordVrcManager *handler.VrcManager
var GlobalTokenManager *handler.TokenManager
var GlobalUserManager *handler.UserManager
var GlobalFileStorage *dao.FileStorage
var GlobalLetterManager *handler.LetterManager
var GlobalMessageManager *handler.MessageManager
var GlobalAttentionManager *handler.AttentionManager