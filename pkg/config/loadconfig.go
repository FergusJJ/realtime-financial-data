package config

import (
	"fmt"
	"os"

	godotenv "github.com/joho/godotenv"
)

func Load(key string) (envVar string, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	envVar = os.Getenv(key)
	if envVar == "" {
		err = fmt.Errorf("%s does not exist", key)
	}
	return envVar, err
}

