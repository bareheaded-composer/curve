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
	if err := m.db.Find(&letters, "sender_uid = ?", senderUID).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return letters, nil
}

func (m *LetterManager) GetHadReceivedLetters(senderUID, receiverUID int) ([]model.Letter, error) {
	letters := make([]model.Letter, 0)
	if err := m.db.Find(&letters, "sender_uid = ? and receiver_uid = ?", senderUID, receiverUID).Error; err != nil {
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
	createTableIfNotExist(m.db, letter, letter.TableName())
	if err := m.db.Create(letter).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
