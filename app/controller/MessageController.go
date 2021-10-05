package controller

import (
	"ajcgo/database/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateMessage(ctx *gin.Context) {
	msg := model.Message{}
	err := ctx.Bind(&msg)

	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(400)
		return
	}

	err = model.CreateMessage(&msg)
	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(500)
		return
	}

	ctx.Status(200)
}

func GetMessage(ctx *gin.Context) {
	msg := *model.GetMessage(ctx.Param("id"))
	log.Debug().Msgf("%+v", msg)

	if &msg == nil {
		ctx.Status(404)
		return
	}

	ctx.JSON(200, msg)
}

func FindMessages(ctx *gin.Context) {
	messages := *model.GetMessagesByUserId(ctx.Param("uid"))

	if len(messages) == 0 {
		ctx.Status(404)
		return
	}

	ctx.JSON(200, messages)
}

func DeleteMessage(ctx *gin.Context) {
	err := model.DeleteMessageById(ctx.Param("id"))
	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(404)
		return
	}
	ctx.Status(204)
}

func DeleteAllMessages(ctx *gin.Context) {
	err := model.DeleteAllMessagesByUserId(ctx.Param("uid"))
	if err != nil {
		log.Warn().Msgf("%v", err)
		ctx.Status(404)
		return
	}
	ctx.Status(204)
}
