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

	err = model.CreateUser(u.Name, u.Email, u.Password)
	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(500)
	}
}

func GetUser(ctx *gin.Context) {
	user := *model.GetUserById(ctx.Param("id"))
	if &user == nil {
		ctx.Status(404)
		return
	}
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

	if &user == nil {
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
		ctx.Status(400)
		return
	}

	err, status := model.UpdateUser(&postedData)
	if err != nil {
		log.Warn().Msgf("%v", err)
	}

	ctx.Status(status)
}

func DeleteUser(ctx *gin.Context) {
	err := model.DeleteUser(ctx.Param("id"))
	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(404)
	}
	ctx.Status(204)
}

func findUsers(ctx *gin.Context) {
	users := *model.GetUsersByName(ctx.Query("name"))

	if len(users) == 0 {
		ctx.Status(404)
		return
	}

	for _, user := range users {
		user.Password = ""
	}

	ctx.JSON(200, users)
}
