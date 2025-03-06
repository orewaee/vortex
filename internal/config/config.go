package config

import (
	"github.com/joho/godotenv"
	"github.com/orewaee/typedenv"
	"os"
	"time"
)

func MustLoad() {
	envs := []string{
		"config/prod.env",
		"config/dev.env",
		"config/postgres.env",
		"config/redis.env",
	}

	for i := range envs {
		err := godotenv.Load(envs[i])
		if err != nil && !os.IsNotExist(err) {
			panic(err)
		}
	}

	typedenv.DefaultString("VORTEX_ADDR", ":8080")
	typedenv.DefaultDuration("ACCESS_LIFETIME", time.Minute*10)
	typedenv.DefaultDuration("REFRESH_LIFETIME", time.Hour*24)
}
