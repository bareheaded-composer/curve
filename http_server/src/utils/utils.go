package utils

import (
	"bytes"
	"curve/src/model"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/disintegration/imaging"
	"image/png"
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

func GetThumbnailDatas(imageDatas [][]byte, newWidth, newHeight int) ([][]byte, error) {
	thumbnailDatas := make([][]byte, 0)
	for _, imageData := range imageDatas {
		thumbnailData, err := GetThumbnailData(imageData, newWidth, newHeight)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		thumbnailDatas = append(thumbnailDatas, thumbnailData)
	}
	return thumbnailDatas, nil
}

func GetThumbnailData(imageData []byte, newWidth, newHeight int) ([]byte, error) {
	image, err := imaging.Decode(bytes.NewBuffer(imageData))
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	image = imaging.Resize(image, newWidth, newHeight, imaging.Lanczos)
	buff := new(bytes.Buffer)
	if err = png.Encode(buff, image); err != nil {
		logs.Error(err)
		return nil, err
	}
	return buff.Bytes(), nil
}
