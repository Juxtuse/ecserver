package config

type RedisCfg struct {
	Addr string
	Pass string
	DB   int
}

type JwtSecrets struct {
	ATKSecret string
	RTKSecret string
}

type S3Config struct {
	Endpoint       string
	PublicEndpoint string
	Key            string
	Secret         string
}

type EnvMode string

const (
	LOCAL   EnvMode = "local"
	DEV     EnvMode = "dev"
	STAGING EnvMode = "staging"
	PROD    EnvMode = "production"
)

func (e EnvMode) IsValid() bool {
	switch e {
	case LOCAL, DEV, STAGING, PROD:
		return true
	default:
		return false
	}
}

type Config struct {
	Env       EnvMode
	DB        string
	Port      string
	Redis     RedisCfg
	JwtSecret JwtSecrets
}
