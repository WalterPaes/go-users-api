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

var (
	recorder = httptest.NewRecorder()

	userCreateInput = dtos.CreateUserInput{
		Name:     "Test",
		Email:    "j@j.com",
		Password: "123456",
	}
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
	type fields struct {
		repository repositories.UserRepository
	}
	type args struct {
		c *gin.Context
		w *httptest.ResponseRecorder
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name: "Success",
			fields: fields{
				repository: repositories.NewUserRepository(mockDatabase(t)),
			},
			args: args{
				c: func(w *httptest.ResponseRecorder) *gin.Context {
					c, _ := gin.CreateTestContext(w)

					jsonBytes, _ := json.Marshal(userCreateInput)

					c.Request = &http.Request{
						Method: http.MethodPost,
						Body:   io.NopCloser(bytes.NewBuffer(jsonBytes)),
					}

					return c
				}(recorder),
				w: recorder,
			},
			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewUserHandler(tt.fields.repository)
			h.CreateUser(tt.args.c)

			assert.Equal(t, tt.wantStatus, tt.args.w.Code)
		})
	}
}
