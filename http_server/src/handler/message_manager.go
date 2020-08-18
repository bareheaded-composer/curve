package handler

import (
	"curve/src/model"
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"net/http"
)

type MessageManager struct {
	db    *gorm.DB
	conns map[int]*websocket.Conn
}

func NewMessageManager(db *gorm.DB) *MessageManager {
	return &MessageManager{
		db:    db,
		conns: make(map[int]*websocket.Conn),
	}
}

func (m *MessageManager) SendMessage(senderUID, receiverUID int, content string) error {
	receiverConn := m.getConn(receiverUID)
	if receiverConn == nil {
		logs.Info("Sending message from sender_uid(%d) to receiver_uid(%d) block, as receiver(%s) is not online", senderUID, receiverUID, receiverUID)
		return nil
	}
	if err := receiverConn.WriteMessage(websocket.TextMessage, []byte(content)); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (m *MessageManager) StoreMessage(senderUID, receiverUID int, content string) error {
	message := &model.Message{
		SenderUID:   senderUID,
		ReceiverUID: receiverUID,
		Content:     content,
	}
	createTableIfNotExist(m.db, message, message.TableName())
	if err := m.db.Create(message).Error; err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (m *MessageManager) SetConn(uid int, c *gin.Context) error {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logs.Error(err)
		return err
	}
	m.conns[uid] = conn
	return nil
}

func (m *MessageManager) getConn(uid int) *websocket.Conn {
	return m.conns[uid]
}
