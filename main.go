package main

import (
	"ajcgo/app/controller"
	"ajcgo/database/model"
	"fmt"
	"github.com/gin-gonic/gin"
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
		user := model.User{Name: "kokkun312312", Email: "oguiku"}
		newUser := model.User{Name: "gami"}

		user.Name = newUser.Name
		user.Email = newUser.Email

		context.JSON(200, user)
	})

	// User
	engine.POST("/api/user", controller.CreateUser)
	engine.GET("/api/user/:id", controller.GetUser)
	engine.GET("/api/user/search", controller.FindUser)
	engine.PATCH("/api/user/:id", controller.UpdateUser)
	engine.DELETE("/api/user/:id", controller.DeleteUser)
	// Message
	engine.POST("/api/message")
	if err = engine.Run(":8080"); err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
