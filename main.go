package main

import (
	"ajcgo/database/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

	engine := gin.Default()

	engine.GET("/test", func(context *gin.Context) {
		context.JSON(200, model.User{
			Model:    gorm.Model{},
			Name:     "Test User",
			Email:    "hello@wor.ld",
			Password: "password",
		})
	})

	if err = engine.Run(":8080"); err != nil {
		log.Fatal().Msgf(
			"%v",
			err,
		)
	}
}
