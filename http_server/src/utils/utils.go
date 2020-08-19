package utils

import (
	"curve/src/model"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/http"
	"strings"
	"time"
)

func IsSizeValid(fileData []byte, maxSize int) bool {
	return GetSize(fileData) <= maxSize
}

func IsTypeValid(fileData []byte, validType []string) bool {
	for _, validType := range validType {
		if GetFileType(fileData) == validType {
			return true
		}
	}
	return false
}

func GetNewFileNameBaseOnTimeFromBase64Data(base64Data string) (string, error) {
	fileData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		logs.Warn(err)
		return model.InvalidFileName, err
	}
	return GetNewFileNameBaseOnTime(fileData), nil
}

func GetNewFileNameBaseOnTime(contentData []byte) string {
	return fmt.Sprintf("%d.%s", time.Now().UnixNano(), GetFileType(contentData))
}

func GetSize(fileData []byte) int {
	return len(fileData)
}

func GetFileType(fileData []byte) string {
	tp := http.DetectContentType(fileData)
	arr := strings.Split(tp, "/")
	if len(arr) != 2 {
		return model.InvalidFileType
	}
	return arr[1]
}
