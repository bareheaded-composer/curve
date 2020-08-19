package utils

import (
	"curve/src/model"
	"encoding/base64"
	"github.com/astaxie/beego/logs"
	"github.com/go-playground/validator/v10"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
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
	for _, validType := range model.ValidAvatarType {
		if GetFileType(fileData) == validType && GetSize(fileData) <= model.AvatarMaxSize {
			return true
		}
	}
	return false
}

func IsValidPhoto(f validator.FieldLevel) bool {
	base64FileData := f.Field().String()
	var fileData []byte
	if _, err := base64.StdEncoding.Decode(fileData, []byte(base64FileData)); err != nil {
		logs.Warn(err)
		return false
	}
	for _, validType := range model.ValidAvatarType {
		if GetFileType(fileData) == validType && GetSize(fileData) <= model.PhotoMaxSize {
			return true
		}
	}
	return false
}
