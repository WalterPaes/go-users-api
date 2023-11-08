package jwt

import (
	"log"
	"time"

	customErrors "github.com/WalterPaes/go-users-api/pkg/errors"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(data map[string]string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(10 * time.Minute).Unix(),
		"sub": data,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", customErrors.New("JWT AUTH ERROR", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Println(customErrors.New("JWT AUTH ERROR", err))
		return false
	}

	return token.Valid
}
