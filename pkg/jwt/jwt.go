package jwt

import (
	"log"
	"time"

	customErrors "github.com/WalterPaes/go-users-api/pkg/errors"
	"github.com/golang-jwt/jwt"
)

type Auth struct {
	secret           []byte
	expiresInMinutes int
}

func NewAuth(secret string, expiresInMinutes int) *Auth {
	return &Auth{
		secret:           []byte(secret),
		expiresInMinutes: expiresInMinutes,
	}
}

func (a *Auth) GenerateToken(data map[string]string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(a.expiresInMinutes) * time.Minute).Unix(),
		"sub": data,
	})

	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		return "", customErrors.New("JWT AUTH ERROR", err)
	}

	return tokenString, nil
}

func (a *Auth) ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return a.secret, nil
	})
	if err != nil {
		log.Println(customErrors.New("JWT AUTH ERROR", err))
		return false
	}

	return token.Valid
}
