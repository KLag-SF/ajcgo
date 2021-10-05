package model

import (
	"ajcgo/database"
	"github.com/rs/zerolog/log"
	"time"
)

type Message struct {
	Id       string    `gorm:"primary key"`
	Uid      string    `gorm:"index:uid"`
	Msg      string    `gorm:"size:140"`
	PostedAt time.Time `gorm:"type:timestamp"`
}

func CreateMessage(msg *Message) error {
	db := database.GetDB()

	msg.Id = getULID()
	msg.PostedAt = time.Now()

	if err := db.Create(msg).Error; err != nil {
		return err
	}

	return nil
}

func GetMessage(id string) *Message {
	db := database.GetDB()

	msg := Message{}
	if err := db.Where("id = ?", id).First(&msg).Error; err != nil {
		log.Warn().Msgf("%v", err)
		return nil
	}

	return &msg
}

func GetMessagesByUserId(uid string) *[]Message {
	db := database.GetDB()

	var messages []Message
	if err := db.Where("uid = ?", uid).Find(&messages).Error; err != nil {
		return nil
	}

	return &messages
}

func DeleteMessageById(id string) error {
	db := database.GetDB()

	if err := db.Delete(&Message{}, id).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAllMessagesByUserId(uid string) error {
	db := database.GetDB()

	if err := db.Delete(&Message{}, "uid = ?", uid).Error; err != nil {
		return err
	}

	return nil
}
