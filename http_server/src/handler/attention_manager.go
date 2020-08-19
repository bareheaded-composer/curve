package handler

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

type AttentionManager struct {
	db *gorm.DB
}

func NewAttentionManager(db *gorm.DB) *AttentionManager {
	return &AttentionManager{
		db: db,
	}
}

func (m *AttentionManager) StoreAttentionIfNotExist(attenderUID, attendeeUID int) error {
	attention := &model.Attention{
		AttenderUID: attenderUID,
		AttendeeUID: attendeeUID,
	}
	createTableIfNotExist(m.db, attention, attention.TableName())
	hasAttended, err := m.HasAttended(attenderUID, attendeeUID)
	if err != nil {
		logs.Error(err)
		return err
	}
	if hasAttended {
		return nil
	}
	if err := m.db.Create(attention).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (m *AttentionManager) HasAttended(attenderUID, attendeeUID int) (bool, error) {
	var attention model.Attention
	if err := m.db.Where("attender_uid = ? and attendee_uid = ?", attenderUID, attendeeUID).First(&attention).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		logs.Error(err)
		return false, err
	}
	return true, nil
}

func (m *AttentionManager) GetAttentionsBaseOnAttendee(attendeeUID int) ([]model.Attention, error) {
	attentions := make([]model.Attention, 0)
	if err := m.db.Where("attendee_uid = ?", attendeeUID).Find(&attentions).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return attentions, nil
}

func (m *AttentionManager) GetAttentionsBaseOnAttender(attenderUID int) ([]model.Attention, error) {
	attentions := make([]model.Attention, 0)
	if err := m.db.Where("attender_uid = ?", attenderUID).Find(&attentions).Error; err != nil {
		logs.Error(err)
		return nil, err
	}
	return attentions, nil
}
