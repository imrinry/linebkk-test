package config

import (
	"line-bk-api/pkg/logs"

	"github.com/joho/godotenv"
)

var (
	DefaultPage  = 1
	DefaultLimit = 10
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logs.Error(err)
	}
}
