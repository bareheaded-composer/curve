package handler

import (
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestVrcFormer(t *testing.T) {
	const charPool = "0123456789"
	const vrcLength = 6
	vrcFormer := NewVrcGenerator(charPool,vrcLength)
	for i:=0;i<10;i++{
		logs.Info(vrcFormer.GetVrc())
	}
}
