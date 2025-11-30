package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Panic().Msgf("missing required environment variable: %s", key)
	}
	return val
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Info().Msg("No .env file found, using default environment variables.")
	}

	cfg := &Config{
		Env:  EnvMode(mustGetEnv("ENV_MODE")),
		DB:   mustGetEnv("DATABASE_URL"),
		Port: mustGetEnv("PORT"),
		Redis: RedisCfg{
			Addr: mustGetEnv("REDIS_ADDR"),
			Pass: os.Getenv("REDIS_PASS"),
			DB:   0,
		},
		JwtSecret: JwtSecrets{
			ATKSecret: mustGetEnv("JWT_ATK_SECRET"),
			RTKSecret: mustGetEnv("JWT_RTK_SECRET"),
		},
	}

	if err := validate(cfg); err != nil {
		log.Panic().Err(err).Msg("Invalid config")
	}

	return cfg
}
