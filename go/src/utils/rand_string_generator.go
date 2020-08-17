package utils

import (
	"math/rand"
)

type RandStringGenerator struct {
	charPool  string
	length int
}

func NewRandStringGenerator(charPool string,length int) *RandStringGenerator {
	return &RandStringGenerator{
		charPool:charPool,
		length:length,
	}
}

func (v *RandStringGenerator) Get() string {
	var str string
	for i := 0; i < v.length; i++ {
		str += string(v.charPool[rand.Intn(len(v.charPool))])
	}
	return str
}
