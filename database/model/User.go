package model

import (
	"ajcgo/database"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func CreateUser(name string, email string, passwd string) {
	db := database.GetDB()
	// Create new user
	user := User{Name: name, Email: email, Password: passwd}
	if err := db.Create(&user).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}
}

func GetUserById(id int) *User {
	db := database.GetDB()

	user := User{}
	if err := db.First(&user, id).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}

	return &user
}

func GetUsersByName(name string) *[]User {
	db := database.GetDB()

	users := []User{}
	if err := db.Where("name = ?", name).Find(&users).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}

	return &users
}

func GetUserByEmail(email string) *User {
	db := database.GetDB()

	user := User{}
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}

	return &user
}

func UpdateUser(id int, name string, email string, passwd string) {
	db := database.GetDB()

	user := User{}
	if err := db.First(&user, id).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}

	user.Name = name
	user.Email = email
	user.Password = passwd
	if err := db.Save(&user).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}
}

func DeleteUser(id int) {
	db := database.GetDB()

	if err := db.Delete(&User{}, id).Error; err != nil {
		log.Warn().Msgf("%v", err)
	}
}
