package utils

import (
	"curve/src/model"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/go-playground/validator/v10"
	"regexp"
)

const (
	VrcLength = 6
	VrcPool   = "0123456789"
)

func IsValidPassword(f validator.FieldLevel) bool {
	const passwordRegexp = `^[_.a-zA-Z0-9]{8,15}$`
	reg := regexp.MustCompile(passwordRegexp)
	return reg.Match([]byte(f.Field().String()))
}

func IsValidVrc(f validator.FieldLevel) bool {
	var vrcRegexp = fmt.Sprintf(`^[%s]{%d,%d}$`, VrcPool, VrcLength, VrcLength)
	reg := regexp.MustCompile(vrcRegexp)
	return reg.Match([]byte(f.Field().String()))
}

func IsValidAvatar(f validator.FieldLevel) bool {
	base64FileData := f.Field().String()
	var fileData []byte
	if _, err := base64.StdEncoding.Decode(fileData, []byte(base64FileData)); err != nil {
		logs.Warn(err)
		return false
	}
	return IsTypeValid(fileData, model.ValidAvatarType) && IsSizeValid(fileData, model.AvatarMaxSize)
}

func IsValidPhoto(f validator.FieldLevel) bool {
	base64FileData := f.Field().String()
	var fileData []byte
	if _, err := base64.StdEncoding.Decode(fileData, []byte(base64FileData)); err != nil {
		logs.Warn(err)
		return false
	}
	return IsTypeValid(fileData, model.ValidPhotoType) && IsSizeValid(fileData, model.PhotoMaxSize)
}

