package model

import (
	"github.com/Roc-zhou/go-util-package/time"
	"gorm.io/gorm"
)

type Model struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreateTime string `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;not null" json:"update_time"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.CreateTime = time.GetNow()
	m.UpdateTime = time.GetNow()
	return nil
}
func (m *Model) BeforeUpdate(tx *gorm.DB) error {
	m.UpdateTime = time.GetNow()
	return nil
}
