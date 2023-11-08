package handlers

import (
	"net/http"

	"github.com/WalterPaes/go-users-api/internal/dtos"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	customErrors "github.com/WalterPaes/go-users-api/pkg/errors"
	jwtAuth "github.com/WalterPaes/go-users-api/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	repository repositories.UserRepository
}

func NewLoginHandler(r repositories.UserRepository) *LoginHandler {
	return &LoginHandler{
		repository: r,
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var loginDto dtos.LoginInput

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.NewValidationErrors(err))
		return
	}

	user, err := h.repository.FindByEmail(loginDto.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	if !user.ValidatePassword(loginDto.Password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwtAuth.GenerateToken(map[string]string{
		"id":       user.ID.String(),
		"username": user.Name,
		"email":    user.Email,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
