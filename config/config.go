package config

import (
	"line-bk-api/pkg/logs"
	"os"

	"github.com/joho/godotenv"
)

var (
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100

	// ประกาศเป็นแค่ตัวแปรเฉยๆ
	JWTSecret           string
	AccessTokenExpired  string
	RefreshTokenExpired string
	JWTIssuer           string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		logs.Error(err)
	}

	// ย้ายการอ่าน env มาไว้ที่นี่
	JWTSecret = os.Getenv("JWT_SECRET")
	AccessTokenExpired = os.Getenv("ACCESS_TOKEN_EXPIRED")
	RefreshTokenExpired = os.Getenv("REFRESH_TOKEN_EXPIRED")
	JWTIssuer = os.Getenv("JWT_ISSUER")
}
