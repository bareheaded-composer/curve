package handler

import (
	"math/rand"
)

type VrcGenerator struct {
	charPool  string
	vrcLength int
}

func NewVrcGenerator(charPool string,vrcLength int) *VrcGenerator{
	return &VrcGenerator{
		charPool:charPool,
		vrcLength:vrcLength,
	}
}

func (v *VrcGenerator) GetVrc() string {
	var vrc string
	for i := 0; i < v.vrcLength; i++ {
		vrc += string(v.charPool[rand.Intn(len(v.charPool))])
	}
	return vrc
}
