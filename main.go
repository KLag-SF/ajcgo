package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}

}
