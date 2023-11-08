package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WalterPaes/go-users-api/internal/dtos"
	"github.com/WalterPaes/go-users-api/internal/entity"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mockDatabase(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	return db
}

func TestUserHandler_CreateUser(t *testing.T) {
	t.Run("Should create an user", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		jsonBytes, _ := json.Marshal(dtos.CreateUserInput{
			Name:     "Test",
			Email:    "j@j.com",
			Password: "123456",
		})

		c.Request = &http.Request{
			Method: http.MethodPost,
			Body:   io.NopCloser(bytes.NewBuffer(jsonBytes)),
		}

		h := NewUserHandler(repositories.NewUserRepository(mockDatabase(t)))
		h.CreateUser(c)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Should has validation errors", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		jsonBytes, _ := json.Marshal(dtos.CreateUserInput{})

		c.Request = &http.Request{
			Method: http.MethodPost,
			Body:   io.NopCloser(bytes.NewBuffer(jsonBytes)),
		}

		h := NewUserHandler(repositories.NewUserRepository(mockDatabase(t)))
		h.CreateUser(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
