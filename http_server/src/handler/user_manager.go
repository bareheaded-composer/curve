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

func (u *UserManager) GetUserInformation(uid int) (*model.UserInformation, error) {
	var userInformation model.UserInformation
	if err := u.db.First(&userInformation, "id = ?", uid).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return &userInformation, nil
}

func (u *UserManager) GetUid(email string) (int, error) {
	var userInformation model.UserInformation
	if err := u.db.First(userInformation, "email = ?", email).Error; err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	return int(userInformation.ID), nil
}

func (u *UserManager) UpdateLastLoginTime(email string) error {
	var userInformation model.UserInformation
	userInformation.LastLoginTime = time.Now()
	if err := u.db.Where("email = ?", email).Table(userInformation.TableName()).Update(&userInformation).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *UserManager) UpdateAvatarFileName(uid int, fileName string) error {
	var userInformation model.UserInformation
	userInformation.AvatarPath = fileName
	if err := u.db.Where("id = ?", uid).Table(userInformation.TableName()).Update(&userInformation).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *UserManager) UpdatePassword(email string, newPassword string) error {
	var userInformation model.UserInformation
	hashSaltyPassword, salt, err := u.getHashSaltyPassword(newPassword)
	if err != nil {
		logs.Error(err)
		return err
	}
	userInformation.Salt = salt
	userInformation.HashSaltyPassword = hashSaltyPassword
	if err := u.db.Where("email = ?", email).Table(userInformation.TableName()).Update(&userInformation).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (u *UserManager) IsPasswordRight(email string, password string) (bool, error) {
	var userInformation model.UserInformation
	if err := u.db.First(&userInformation, "email = ?", email).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		logs.Error(err)
		return false, err
	}
	saltyPassword := getSaltyPassword(password, userInformation.Salt)
	hashSaltyPassword, err := u.hasher.GetHashString(saltyPassword)
	if err != nil {
		logs.Error(err)
		return false, err
	}
	return hashSaltyPassword == userInformation.HashSaltyPassword, nil
}

func (u *UserManager) IsEmailExist(email string) (bool, error) {
	var userInformation model.UserInformation
	if !u.db.HasTable(userInformation) {
		return false, nil
	}
	if err := u.db.First(&userInformation, "email = ?", email).Error; err != nil {
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
	userInformation := &model.UserInformation{
		Email:             email,
		HashSaltyPassword: hashSaltyPassword,
		Type:              tp,
		LastLoginTime:     time.Now(),
		Salt:              salt,
		AvatarPath:        model.DefaultAvatarPath,
		Username:          model.DefaultUsername,
		Sex:               model.DefaultSex,
		ContactPhone:      model.DefaultContactPhone,
		ContactEmail:      model.DefaultContactEmail,
		Birthday:          time.Now(),
	}
	createTableIfNotExist(u.db, userInformation, userInformation.TableName())
	if err := u.db.Save(userInformation).Error; err != nil {
		logs.Error(err)
		return model.FlagOfInvalidUID, err
	}
	return int(userInformation.ID), nil
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
