package Middleware

import (
	"ajcgo/app/auth"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"github.com/rs/zerolog/log"
)

func LoginCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		loginUserJson, err := dproxy.New(session.Get("loginUser")).String()

		if err != nil {
			log.Warn().Msgf("%v", err)
			ctx.Status(401)
			ctx.Abort()
			return
		}

		var loginInfo auth.AuthUser
		err = json.Unmarshal([]byte(loginUserJson), &loginInfo)
	}
}
