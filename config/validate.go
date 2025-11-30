package config

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	errInvalidDBUrl      = errors.New("invalid DB URL passed on config")
	errRedisPassNotEmpty = errors.New("redis password can't be empty string on production")
	errInvalidModes      = errors.New("MODE should be 'LOCAL' | 'DEV' | 'STAGING' | 'PRODUCTION'")
	errJwtInvalidLength  = errors.New("JWT secrets should be minimum 45 characters on production")
)

func validate(cfg *Config) error {
	if !cfg.Env.IsValid() {
		return errInvalidModes
	}

	_, err := pgx.ParseConfig(cfg.DB)
	if err != nil {
		return errInvalidDBUrl
	}

	if cfg.Env == "production" {
		if cfg.Redis.Pass == "" {
			return errRedisPassNotEmpty
		}

		if len(cfg.JwtSecret.ATKSecret) > 44 || len(cfg.JwtSecret.RTKSecret) > 44 {
			return errJwtInvalidLength
		}
	}

	return nil
}
