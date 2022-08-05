package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/swaingotnochill/go-csrf-template/utils"
)

// ------------------------------------------------------------

type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

// ------------------------------------------------------------

type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

// ------------------------------------------------------------

const RefreshTokenExpiration = time.Hour * 72
const AuthTokenExpiration = time.Minute * 15

// ------------------------------------------------------------

func GenerateCSRFSecret() (string, error) {
	return utils.GenerateRandomStrings(32)
}
