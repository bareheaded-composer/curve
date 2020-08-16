package utils

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestCoder(t *testing.T) {
	coder := NewCoder("1234567891234567")
	originStr := "i love you"
	encryptStr, err := coder.Encrypt(originStr)
	if err != nil {
		logs.Error(encryptStr)
		return
	}
	decryptStr, err := coder.Decrypt(encryptStr)
	if err != nil {
		logs.Error(encryptStr)
		return
	}
	if decryptStr != originStr {
		logs.Error("Test failed. %s(decryptStr) != %s(originStr)", decryptStr, originStr)
	} else {
		logs.Info("Test passed. %s(decryptStr) == %s(originStr)", decryptStr, originStr)
	}
}
