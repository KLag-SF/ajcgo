package controller

import (
	"ajcgo/database/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateUser(ctx *gin.Context) {
	u := model.User{}
	err := ctx.Bind(&u)

	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(400)
		return
	}

	model.CreateUser(u.Name, u.Email, u.Password)
}

func GetUser(ctx *gin.Context) {
	user := *model.GetUserById(ctx.Param("id"))
	user.Password = ""
	ctx.JSON(200, user)
}

func FindUser(ctx *gin.Context) {
	user := model.User{}

	if ctx.Query("email") != "" {
		user = *model.GetUserByEmail(ctx.Query("email"))
	} else if ctx.Query("name") != "" {
		findUsers(ctx)
	} else {
		ctx.Status(404)
		return
	}

	// Remove password hash
	user.Password = ""
	ctx.JSON(200, user)
}

func UpdateUser(ctx *gin.Context) {
	postedData := model.User{ID: ctx.Param("id")}

	if err := ctx.Bind(&postedData); err != nil {
		log.Warn().Msgf("%v", err)
		return
	}

	model.UpdateUser(&postedData)
}

func DeleteUser(ctx *gin.Context) {
	model.DeleteUser(ctx.Param("id"))
}

func findUsers(ctx *gin.Context) {
	users := *model.GetUsersByName(ctx.Query("name"))
	for _, user := range users {
		user.Password = ""
	}
	ctx.JSON(200, users)
}
