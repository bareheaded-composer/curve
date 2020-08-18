package handler

import (
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

func createTableIfNotExist(db *gorm.DB,table interface{},tableName string) {
	if !db.HasTable(table) {
		logs.Info("As table(%s) not exist, it will be created.", tableName)
		db.CreateTable(table)
		logs.Info("Creating table(%s) success.", tableName)
	}
}
