package utils

import (
	"crypto/sha1"
	"curve/src/model"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type Hasher struct {
	hashKey string
}

func NewHasher(hashKey string) *Hasher{
	return &Hasher{
		hashKey: hashKey,
	}
}
func (h *Hasher) GetHashString(originString string) (string, error) {
	shaer := sha1.New()
	if _, err := shaer.Write([]byte(originString)); err != nil {
		logs.Error(err)
		return model.InvalidHashString, err
	}
	return fmt.Sprintf("%x", shaer.Sum(nil)), nil
}
