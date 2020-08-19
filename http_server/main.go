package main

import (
	"curve/src/controller"
	"curve/src/dao"
	"curve/src/env"
	"curve/src/handler"
	"curve/src/model"
	"curve/src/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"html/template"
	"time"
)

const (
	confPath = "conf/conf.json"
)

func init() {
	InitConf()
	InitLogger()
	InitGlobalRegisterVrcManager()
	InitGlobalChangePasswordVrcManager()
	InitGlobalTokenManager()
	InitGlobalUserManager()
	InitGlobalLetterManager()
	InitGlobalFileStorage()
	InitValidatorEngine()
}

func InitLogger() {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
}
func InitConf() {
	if err := env.Conf.Load(confPath); err != nil {
		logs.Error("Load conf(%s) failed.", confPath)
		return
	}
}
func InitGlobalRegisterVrcManager() {
	vrcGenerator := utils.NewRandStringGenerator(
		env.Conf.Vrc.CharPool, env.Conf.Vrc.Length,
	)
	client := utils.NewEmailClient(
		env.Conf.EmailClient.EmailAddr,
		env.Conf.EmailClient.AuthCode,
		env.Conf.EmailClient.SmtpAddr,
		env.Conf.EmailClient.SmtpPort,
	)
	emailTemplate := template.New("")
	if _, err := emailTemplate.Parse(env.Conf.Template.RegisterEmailTemplate); err != nil {
		logs.Error(err)
		return
	}
	vrcEmailSender := utils.NewVrcEmailSender(client, vrcGenerator, emailTemplate)
	cache := dao.NewCache(
		env.Conf.Cache.Network,
		env.Conf.Cache.Host,
	)
	controller.GlobalRegisterVrcManager = handler.NewVrcManager(
		vrcEmailSender,
		cache,
		env.Conf.Vrc.ExpiredSecond,
		"注册邮件",
		"register",
	)
}
func InitGlobalChangePasswordVrcManager() {
	vrcGenerator := utils.NewRandStringGenerator(
		env.Conf.Vrc.CharPool, env.Conf.Vrc.Length,
	)
	client := utils.NewEmailClient(
		env.Conf.EmailClient.EmailAddr,
		env.Conf.EmailClient.AuthCode,
		env.Conf.EmailClient.SmtpAddr,
		env.Conf.EmailClient.SmtpPort,
	)
	emailTemplate := template.New("")
	if _, err := emailTemplate.Parse(env.Conf.Template.ChangePasswordEmailTemplate); err != nil {
		logs.Error(err)
		return
	}
	vrcEmailSender := utils.NewVrcEmailSender(client, vrcGenerator, emailTemplate)
	cache := dao.NewCache(
		env.Conf.Cache.Network,
		env.Conf.Cache.Host,
	)
	controller.GlobalChangePasswordVrcManager = handler.NewVrcManager(
		vrcEmailSender,
		cache,
		env.Conf.Vrc.ExpiredSecond,
		"修改密码邮件",
		"changePassword",
	)
}
func InitGlobalTokenManager() {
	coder := utils.NewCoder(env.Conf.SecretKey.ForCoder)
	tokenDuration := 72 * time.Hour
	controller.GlobalTokenManager = handler.NewTokenManager(coder, tokenDuration, env.Conf.SecretKey.ForToken, model.KeyForUid)
}
func InitGlobalUserManager() {
	saltGenerator := utils.NewRandStringGenerator(env.Conf.Salt.CharPool, env.Conf.Salt.Length)
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.Conf.Mysql.User,
			env.Conf.Mysql.Password,
			env.Conf.Mysql.Host,
			env.Conf.Mysql.DBName,
		),
	)
	if err != nil {
		logs.Error(err)
		return
	}
	hasher := utils.NewHasher(env.Conf.SecretKey.ForHasher)
	controller.GlobalUserManager = handler.NewUserManager(db, saltGenerator, hasher)
}
func InitGlobalLetterManager() {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			env.Conf.Mysql.User,
			env.Conf.Mysql.Password,
			env.Conf.Mysql.Host,
			env.Conf.Mysql.DBName,
		),
	)
	if err != nil {
		logs.Error(err)
		return
	}
	controller.GlobalLetterManager = handler.NewLetterManager(db)
}
func InitGlobalFileStorage() {
	const rootPath = "../assert"
	controller.GlobalFileStorage = dao.NewFileStorage(rootPath)
}
func InitValidatorEngine() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("password", utils.IsValidPassword); err != nil {
			logs.Error(err)
			return
		}
		if err := v.RegisterValidation("vrc", utils.IsValidVrc); err != nil {
			logs.Error(err)
			return
		}
		if err := v.RegisterValidation("avatar", utils.IsValidAvatar); err != nil {
			logs.Error(err)
			return
		}
		if err := v.RegisterValidation("photo", utils.IsValidPhoto); err != nil {
			logs.Error(err)
			return
		}
	}
}


func main() {
	r := gin.Default()
	r.GET("/test", controller.Test)

	publicGroup := r.Group("/v1/public")
	{
		publicGroup.GET("/avatar/:name", controller.Avatar)
		publicGroup.GET("/photo", controller.Photo)
	}

	touristGroup := r.Group("/v1/tourist")
	{
		touristGroup.PATCH("/password", controller.ChangePassword)
		touristGroup.POST("/ask_for_register", controller.AskForRegister)
		touristGroup.POST("/ask_for_change_password", controller.AskForChangePassword)
		touristGroup.POST("/login", controller.Login)
		touristGroup.POST("/user", controller.Register)
	}

	userGroup := r.Group("/v1/user")
	{
		userGroup.PATCH("/avatar", controller.UpdateAvatar)
		userGroup.POST("/photo", controller.UpLoadPhoto)
		userGroup.POST("/letter", controller.SendLetter)
		userGroup.POST("/message", controller.SendMessage)
		userGroup.POST("/receiving_message_client", controller.RegisterClientOfReceivingMessage)
		userGroup.GET("/had_sent_letter", controller.HadSentLetter)
		userGroup.GET("/had_received_letter/:sender_uid", controller.HadReceivedLetter)
	}

	if err := r.Run(fmt.Sprintf(":%d", env.Conf.Http.Port)); err != nil {
		logs.Error("Running go http server failed. :|")
		return
	}
}
