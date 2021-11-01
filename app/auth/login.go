package auth

import (
	"ajcgo/database/model"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.BindJSON(&req)

	if err != nil {
		log.Warn().Msgf("failed to bind JSON:%v", err)
		ctx.Status(400)
		return
	}

	user := *model.GetUserByEmail(req.Email)
	if user.ID == "" {
		log.Warn().Msg("User not found.")
		ctx.Status(400)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		log.Warn().Msgf("main.go:30:%v", err)
		ctx.Status(401)
		return
	}

	session := sessions.Default(ctx)
	loginUser, err := json.Marshal(user)

	if err == nil {
		session.Set("loginUser", string(loginUser))
		err = session.Save()
		log.Debug().Msgf("%+v", err)
		if err != nil {
			log.Warn().Msgf("%v", err)
			ctx.Status(400)
			return
		}

		ctx.Redirect(301, "/auth")
	} else {
		ctx.Status(500)
	}
}
