package config

import (
	"github.com/joho/godotenv"
	"github.com/orewaee/typedenv"
	"os"
	"time"
)

func MustLoad() {
	err := godotenv.Load("config/.env")

	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	typedenv.DefaultString("VORTEX_ADDR", ":8080")

	typedenv.DefaultDuration("ACCESS_LIFETIME", time.Minute*10)
	typedenv.DefaultDuration("REFRESH_LIFETIME", time.Hour*24)

	typedenv.DefaultString("ALPHABET", "abcdefghijklmnopqrstuvwxyz")

	typedenv.DefaultString("SUPER_ID", "00000000")
}
