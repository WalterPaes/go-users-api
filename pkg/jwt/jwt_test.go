package jwt_test

import (
	"testing"

	"github.com/WalterPaes/go-users-api/pkg/jwt"
	"github.com/stretchr/testify/assert"
)

func TestJwtAuth(t *testing.T) {
	token, err := jwt.GenerateToken(map[string]string{
		"name":  "Teste",
		"email": "test@email.com",
	})

	if err != nil {
		t.Error(err)
	}

	isValid := jwt.ValidateToken(token)

	assert.True(t, isValid)
}
