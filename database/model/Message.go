package model

import (
	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	UserId  string `gorm:"index:user_id"`
	Message string `gorm:"size:140"`
}
