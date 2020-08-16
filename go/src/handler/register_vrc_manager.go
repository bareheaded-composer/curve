package handler

import (
	"curve/src/dao"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type RegisterVrcManager struct {
	vrcEmailSender   *VrcEmailSender
	vrcStorage       *dao.Cache
	vrcExpiredSecond int
	vrcEmailSubject string
}

func NewRegisterVrcManager(vrcEmailSender *VrcEmailSender, vrcStorage *dao.Cache, vrcExpiredSecond int) *RegisterVrcManager {
	return &RegisterVrcManager{
		vrcEmailSender:   vrcEmailSender,
		vrcStorage:       vrcStorage,
		vrcExpiredSecond: vrcExpiredSecond,
		vrcEmailSubject:"验证码邮件",
	}
}

func (r *RegisterVrcManager) SendAndStoreVrc(registerEmail string) error {
	vrc, err := r.vrcEmailSender.SendVrcEmail(r.vrcEmailSubject, registerEmail,r.vrcExpiredSecond)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := r.vrcStorage.Set(r.getStoreKey(registerEmail), []byte(vrc), r.vrcExpiredSecond); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (r *RegisterVrcManager) IsVrcRight(registerEmail string, checkVrc string) (bool, error) {
	rightVrcBytes, err := r.vrcStorage.Get(r.getStoreKey(registerEmail))
	if err != nil {
		logs.Error(err)
		return false, err
	}
	return string(rightVrcBytes) == checkVrc, nil
}

func (r *RegisterVrcManager) DelVrcOfRegisterEmail(registerEmail string) error {
	if err := r.vrcStorage.Del(r.getStoreKey(registerEmail)); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (r *RegisterVrcManager) getStoreKey(registerEmail string) string {
	return fmt.Sprintf("%s:vrc", registerEmail)
}
