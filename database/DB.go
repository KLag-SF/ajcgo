package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
	"os"
)

func GetDB() *gorm.DB {
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := fmt.Sprintf("tcp(%v)", os.Getenv("DB_ADDR"))
	DBNAME := os.Getenv("DB")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	return db
}
