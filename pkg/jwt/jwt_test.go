package jwt_test

import (
	"testing"

	"github.com/WalterPaes/go-users-api/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuth(t *testing.T) {
	jwtAuth := jwt.NewAuth("secret", 10)

	token, err := jwtAuth.GenerateToken(map[string]string{
		"name":  "Teste",
		"email": "test@email.com",
	})

	if err != nil {
		t.Error(err)
	}

	isValid := jwtAuth.ValidateToken(token)

	assert.True(t, isValid)
}
