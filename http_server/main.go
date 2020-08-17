package main

import (
	"curve/src/controller"
	"curve/src/dao"
	"curve/src/env"
	"curve/src/handler"
	"curve/src/model"
	"curve/src/utils"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"html/template"
	"regexp"
	"time"
)

const (
	VrcLength = 6
	VrcPool   = "0123456789"
	confPath  = "conf/conf.json"
)

func init() {
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
}

func init() {
	if err := env.Conf.Load(confPath); err != nil {
		logs.Error("Load conf(%s) failed.", confPath)
		return
	}
}

func init() {
	const (
		charPool             = "0123456789"
		vrcLength            = 6
		vrcExpiredSecond     = 60
		emailTemplateContent = `您的验证码是: {{.Vrc}} 验证码过期时间为: {{.VrcExpiredSecond}}s.`
	)
	vrcGenerator := utils.NewRandStringGenerator(charPool, vrcLength)
	client := utils.NewEmailClient(
		env.Conf.EmailClient.EmailAddr,
		env.Conf.EmailClient.AuthCode,
		env.Conf.EmailClient.SmtpAddr,
		env.Conf.EmailClient.SmtpPort,
	)
	emailTemplate := template.New("")
	if _, err := emailTemplate.Parse(emailTemplateContent); err != nil {
		logs.Error(err)
		return
	}
	vrcEmailSender := handler.NewVrcEmailSender(client, vrcGenerator, emailTemplate)
	cache := dao.NewCache(
		env.Conf.Cache.Network,
		env.Conf.Cache.Host,
	)
	controller.GlobalRegisterVrcManager = handler.NewVrcManager(
		vrcEmailSender,
		cache,
		vrcExpiredSecond,
		"注册邮件",
		"register",
	)
}

func init() {
	const (
		charPool             = "0123456789"
		vrcLength            = 6
		vrcExpiredSecond     = 60
		emailTemplateContent = `您的验证码是: {{.Vrc}} 验证码过期时间为: {{.VrcExpiredSecond}}s.`
	)
	vrcGenerator := utils.NewRandStringGenerator(charPool, vrcLength)
	client := utils.NewEmailClient(
		env.Conf.EmailClient.EmailAddr,
		env.Conf.EmailClient.AuthCode,
		env.Conf.EmailClient.SmtpAddr,
		env.Conf.EmailClient.SmtpPort,
	)
	emailTemplate := template.New("")
	if _, err := emailTemplate.Parse(emailTemplateContent); err != nil {
		logs.Error(err)
		return
	}
	vrcEmailSender := handler.NewVrcEmailSender(client, vrcGenerator, emailTemplate)
	cache := dao.NewCache(
		env.Conf.Cache.Network,
		env.Conf.Cache.Host,
	)
	controller.GlobalChangePasswordVrcManager = handler.NewVrcManager(
		vrcEmailSender,
		cache,
		vrcExpiredSecond,
		"修改密码邮件",
		"changePassword",
	)
}

func init() {
	const coderSecretKey = "1234567891234567"
	const tokenSecretKey = "abcdefghi"
	coder := utils.NewCoder(coderSecretKey)
	tokenDuration := 72 * time.Hour
	controller.GlobalTokenAnnouncer = handler.NewTokenAnnouncer(coder, tokenDuration, tokenSecretKey, model.KeyForUid)
}

func init() {
	const (
		charPool  = "0123456789"
		vrcLength = 6
	)
	saltGenerator := utils.NewRandStringGenerator(charPool, vrcLength)
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
	const hashKey = "12345678"
	hasher := utils.NewHasher(hashKey)
	controller.GlobalUserManager = handler.NewUserManager(db, saltGenerator, hasher)
}

func init() {
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

func init() {
	const rootPath = "../assert"
	controller.GlobalFileStorage = dao.NewFileStorage(rootPath)
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("password", password); err != nil {
			logs.Error(err)
			return
		}
		if err := v.RegisterValidation("vrc", vrc); err != nil {
			logs.Error(err)
			return
		}
		if err := v.RegisterValidation("avatar", avatar); err != nil {
			logs.Error(err)
			return
		}
	}
	r.GET("/test", controller.Test)
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.POST("/askForRegister", controller.AskForRegister)
	r.POST("/askForChangePassword", controller.AskForChangePassword)
	r.POST("/changePassword", controller.ChangePassword)
	r.POST("/updateAvatar", controller.UpdateAvatar)
	r.POST("/sendLetter", controller.SendLetter)
	r.GET("/avatar/:name", controller.Avatar)
	r.GET("/hadSentLetter",controller.HadSentLetter)
	r.GET("/hadReceivedLetter/:senderUID",controller.HadReceivedLetter)
	if err := r.Run(fmt.Sprintf(":%d", env.Conf.Http.Port)); err != nil {
		logs.Error("Running go http server failed. :|")
		return
	}
}

func password(f validator.FieldLevel) bool {
	const passwordRegexp = `^[_.a-zA-Z0-9]{8,15}$`
	reg := regexp.MustCompile(passwordRegexp)
	return reg.Match([]byte(f.Field().String()))
}

func vrc(f validator.FieldLevel) bool {
	var vrcRegexp = fmt.Sprintf(`^[%s]{%d,%d}$`, VrcPool, VrcLength, VrcLength)
	reg := regexp.MustCompile(vrcRegexp)
	return reg.Match([]byte(f.Field().String()))
}

func avatar(f validator.FieldLevel) bool {
	base64FileData := f.Field().String()
	var fileData []byte
	if _, err := base64.StdEncoding.Decode(fileData, []byte(base64FileData)); err != nil {
		logs.Warn(err)
		return false
	}
	for _, validType := range model.ValidAvatarType {
		if utils.GetFileType(fileData) == validType && utils.GetSize(fileData) <= model.AvatarMaxSize {
			return true
		}
	}
	return false
}
