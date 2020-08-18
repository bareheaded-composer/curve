package handler

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type LetterManager struct {
	db *gorm.DB
}

func NewLetterManager(db *gorm.DB) *LetterManager {
	return &LetterManager{
		db: db,
	}
}

func (m *LetterManager) GetHadSentLetters(senderUID int) ([]model.Letter, error) {
	letters := make([]model.Letter, 0)
	if err := m.db.Where("sender_uid = ?", senderUID).Find(&letters).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return letters, nil
}

func (m *LetterManager) GetHadReceivedLetters(senderUID, receiverUID int) ([]model.Letter, error) {
	letters := make([]model.Letter, 0)
	if err := m.db.Where("sender_uid = ? and receiver_uid = ?", senderUID, receiverUID).Find(&letters).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return letters, nil
}

func (m *LetterManager) StoreLetter(senderUID, receiverUID int, content string) error {
	letter := &model.Letter{
		SenderUID:   senderUID,
		ReceiverUID: receiverUID,
		Content:     content,
	}
	if !m.db.HasTable(letter) {
		logs.Info("As table(%s) not exist, it will be created.", letter.TableName())
		m.db.CreateTable(letter)
		logs.Info("Creating table(%s) success.", letter.TableName())
	}
	if err := m.db.Create(letter).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
