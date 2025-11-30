package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Info().Msg("No .env file found, using default environment variables.")
	}

}
