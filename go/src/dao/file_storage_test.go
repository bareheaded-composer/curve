package dao

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestFileStorage(t *testing.T) {
	photoStorage := NewFileStorage(".")
	logs.Info(photoStorage.Store("hello.jpg",[]byte{1,2,3,4}))
	logs.Info(photoStorage.Get("hello.jpg"))
}
