package model

import (
	"ajcgo/database"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type Message struct {
	gorm.Model
	UserId string `gorm:"index:user_id"`
	Text   string `gorm:"size:140"`
}

func CreateMessage(txt string) {
	db := database.GetDB()

	m := Message{Text: txt}
	if err := db.Create(&m).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}
}

func GetMessage(id int) *Message {
	db := database.GetDB()

	msg := Message{}
	if err := db.First(&msg, id).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}

	return &msg
}

func DeleteMessage(id int) {
	db := database.GetDB()

	if err := db.Delete(&Message{}, id).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}
}
