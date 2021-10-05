package model

import (
	"ajcgo/database"
	"github.com/oklog/ulid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type User struct {
	ID       string `gorm:"primary_key"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func CreateUser(name string, email string, passwd string) error {
	db := database.GetDB()

	// Hash password
	bytePW := []byte(passwd)
	hashedPW, _ := bcrypt.GenerateFromPassword(bytePW, 10)

	// Create new user
	user := User{Name: name, Email: email, Password: string(hashedPW)}
	user.ID = getULID()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserById(id string) *User {
	db := database.GetDB()

	user := User{}
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Warn().Msgf("%v", err)
		return nil
	}

	return &user
}

func GetUsersByName(name string) *[]User {
	db := database.GetDB()

	var users []User
	if err := db.Where("name = ?", name).Find(&users).Error; err != nil {
		log.Warn().Msgf("%v", err)
		return nil
	}

	return &users
}

func GetUserByEmail(email string) *User {
	db := database.GetDB()

	user := User{}
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Warn().Msgf("%v", err)
		return nil
	}

	return &user
}

func UpdateUser(u *User) (error, int) {
	db := database.GetDB()
	user := GetUserById(u.ID)

	if &user == nil {
		return nil, 404
	}

	if u.Name != "" {
		user.Name = u.Name
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Password != "" {
		bytePW := []byte(u.Password)
		hashedPW, _ := bcrypt.GenerateFromPassword(bytePW, 10)
		user.Password = string(hashedPW)
	}

	log.Debug().Msgf("%+v", user)

	if err := db.Save(&user).Error; err != nil {
		return err, 500
	}

	return nil, 200
}

func DeleteUser(id string) error {
	db := database.GetDB()

	if err := db.Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func getULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
