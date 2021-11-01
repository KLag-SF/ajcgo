package main

import (
	"ajcgo/app/Middleware"
	"ajcgo/app/auth"
	"ajcgo/app/controller"
	"ajcgo/database/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	engine.LoadHTMLGlob("resources/templates/*.html")

	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mySession", store))

	engine.GET("/test", func(context *gin.Context) {
		user := model.User{Name: "kokkun312312", Email: "oguiku"}
		newUser := model.User{Name: "gami"}

		user.Name = newUser.Name
		user.Email = newUser.Email

		context.JSON(200, user)
	})

	engine.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	// User
	engine.POST("/api/user", controller.CreateUser)
	engine.GET("/api/user/:id", controller.GetUser)
	engine.GET("/api/user/search", controller.FindUser)
	engine.PATCH("/api/user/:id", controller.UpdateUser)
	engine.DELETE("/api/user/:id", controller.DeleteUser)
	// Message
	engine.POST("/api/message", controller.CreateMessage)
	engine.GET("/api/message/:id", controller.GetMessage)
	engine.GET("/api/message/user/:uid", controller.FindMessages)
	engine.DELETE("/api/message/:id", controller.DeleteMessage)
	engine.DELETE("/api/message/user/:uid", controller.DeleteAllMessages)
	// Login
	engine.GET("/login", func(context *gin.Context) {
		context.Status(404)
	})
	engine.POST("/login", auth.Login)
	// Authenticated user only
	authUserGroup := engine.Group("/auth")
	authUserGroup.Use(Middleware.LoginCheck())
	{
		authUserGroup.GET("/", func(context *gin.Context) {
			type test struct {
				message string
			}
			msg := test{message: "You're logged in!"}
			context.JSON(200, msg)
		})

	}
	if err = engine.Run(":8080"); err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
