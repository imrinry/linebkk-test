package utils

import (
	"errors"
	"line-bk-api/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(userID string) (string, error) {
	jwtExpired, err := time.ParseDuration(config.AccessTokenExpired)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(jwtExpired).Unix(),
		"iat": time.Now().Unix(),
		"iss": config.JWTIssuer,
	})
	return token.SignedString([]byte(config.JWTSecret))
}

func GenerateRefreshToken(userID string) (string, error) {
	jwtExpired, err := time.ParseDuration(config.RefreshTokenExpired)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(jwtExpired).Unix(),
		"iat": time.Now().Unix(),
		"iss": config.JWTIssuer,
	})
	return token.SignedString([]byte(config.JWTSecret))
}

func ValidateAccessToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims)["sub"].(string), nil
}
