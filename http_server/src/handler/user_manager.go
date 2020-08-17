package handler

import (
	"curve/src/model"
	"curve/src/utils"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"time"
)

type UserManager struct {
	db            *gorm.DB
	saltGenerator *utils.RandStringGenerator
	hasher        *utils.Hasher
}

func NewUserManager(db *gorm.DB, saltGenerator *utils.RandStringGenerator, hasher *utils.Hasher) *UserManager {
	return &UserManager{
		db:            db,
		saltGenerator: saltGenerator,
		hasher:        hasher,
	}
}

func (u *UserManager) GetUpi(uid int) (*model.UPI, error) {
	upi := &model.UPI{}
	if err := u.db.Where("id = ?", uid).First(upi).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return upi, nil
}

func (u *UserManager) GetUid(email string) (int, error) {
	uai := &model.UAI{}
	if err := u.db.Where("email = ?", email).First(uai).Error; err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	return int(uai.ID), nil
}

func (u *UserManager) UpdateLastLoginTime(email string) error {
	uai := &model.UAI{}
	uai.LastLoginTime = time.Now()
	if err := u.db.Where("email = ?", email).Table(uai.TableName()).Update(uai).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *UserManager) UpdateAvatarFileName(uid int, fileName string) error {
	upi := &model.UPI{}
	upi.AvatarPath = fileName
	if err := u.db.Where("id = ?", uid).Table(upi.TableName()).Update(upi).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *UserManager) UpdatePassword(email string, newPassword string) error {
	uai := &model.UAI{}
	hashSaltyPassword, salt, err := u.getHashSaltyPassword(newPassword)
	if err != nil {
		logs.Error(err)
		return err
	}
	uai.Salt = salt
	uai.HashSaltyPassword = hashSaltyPassword
	if err := u.db.Where("email = ?", email).Table(uai.TableName()).Update(uai).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *UserManager) IsPasswordRight(email string, password string) (bool, error) {
	uai := &model.UAI{}
	err := u.db.Where("email = ?", email).First(uai).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		logs.Error(err)
		return false, err
	}
	saltyPassword := getSaltyPassword(password, uai.Salt)
	hashSaltyPassword, err := u.hasher.GetHashString(saltyPassword)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	return hashSaltyPassword == uai.HashSaltyPassword, nil
}

func (u *UserManager) IsEmailExist(email string) (bool, error) {
	uai := &model.UAI{}
	err := u.db.Where("email = ?", email).First(uai).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		logs.Error(err)
		return false, err
	}
	return true, nil
}

func (u *UserManager) InsertUser(email, password string, tp model.UserType) (int, error) {
	hashSaltyPassword, salt, err := u.getHashSaltyPassword(password)
	if err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	uai := &model.UAI{
		Email:             email,
		HashSaltyPassword: hashSaltyPassword,
		Type:              tp,
		LastLoginTime:     time.Now(),
		Salt:              salt,
	}
	upi := &model.UPI{
		AvatarPath:   model.DefaultAvatarPath,
		Username:     model.DefaultUsername,
		Sex:          model.DefaultSex,
		ContactPhone: model.DefaultContactPhone,
		ContactEmail: model.DefaultContactEmail,
		Birthday:     time.Now(),
	}
	if !u.db.HasTable(uai) {
		logs.Info("As table(%s) not exist, it will be created.", uai.TableName())
		u.db.CreateTable(uai)
		logs.Info("Creating table(%s) success.", uai.TableName())
	}
	if !u.db.HasTable(upi) {
		logs.Info("As table(%s) not exist, it will be created.", upi.TableName())
		u.db.CreateTable(upi)
		logs.Info("Creating table(%s) success.", upi.TableName())
	}
	if err := u.db.Save(uai).Error; err != nil {
		logs.Info("Inserting User(%s)'s UAI fail.", email)
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	upi.ID = uai.ID
	logs.Info("Inserting User(%s)'s UAI success.", email)
	if err := u.db.Save(upi).Error; err != nil {
		logs.Info("Inserting User(%s)'s UPI fail.", email)
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	logs.Info("Inserting User(%s)'s UPI success.", email)
	return int(uai.ID), nil
}

func (u *UserManager) getHashSaltyPassword(password string) (string, string, error) {
	salt := u.saltGenerator.Get()
	saltyPassword := getSaltyPassword(password, salt)
	hashSaltyPassword, err := u.hasher.GetHashString(saltyPassword)
	if err != nil {
		logs.Error(err)
		return model.InvalidHashString, model.InvalidSalt, err
	}
	return hashSaltyPassword, salt, nil
}

func getSaltyPassword(password, salt string) string {
	return password + salt
}
