package handler

import (
	"curve/src/dao"
	"curve/src/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type VrcManager struct {
	vrcEmailSender   *utils.VrcEmailSender
	vrcEmailSubject  string
	vrcStorage       *dao.Cache
	vrcExpiredSecond int
	keyPrefix        string
}

func NewVrcManager(
	vrcEmailSender *utils.VrcEmailSender,
	vrcStorage *dao.Cache,
	vrcExpiredSecond int,
	vrcEmailSubject string,
	keyPrefix string,
) *VrcManager {
	return &VrcManager{
		vrcStorage:       vrcStorage,
		vrcExpiredSecond: vrcExpiredSecond,
		keyPrefix:        keyPrefix,
		vrcEmailSubject:  vrcEmailSubject,
		vrcEmailSender:   vrcEmailSender,
	}
}

func (r *VrcManager) SendAndStoreVrc(email string) error {
	vrc, err := r.vrcEmailSender.SendVrcEmail(r.vrcEmailSubject, email, r.vrcExpiredSecond)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := r.vrcStorage.Set(r.getStoreKey(email), []byte(vrc), r.vrcExpiredSecond); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (r *VrcManager) IsVrcRight(email string, checkVrc string) (bool, error) {
	rightVrcBytes, err := r.vrcStorage.Get(r.getStoreKey(email))
	if err != nil {
		logs.Error(err)
		return false, err
	}
	return string(rightVrcBytes) == checkVrc, nil
}

func (r *VrcManager) DelVrc(email string) error {
	if err := r.vrcStorage.Del(r.getStoreKey(email)); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (r *VrcManager) getStoreKey(email string) string {
	return fmt.Sprintf("%s:%s", r.keyPrefix, email)
}
