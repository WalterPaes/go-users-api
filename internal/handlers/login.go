package handlers

import (
	"net/http"

	"github.com/WalterPaes/go-users-api/internal/dtos"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	customErrors "github.com/WalterPaes/go-users-api/pkg/errors"
	"github.com/WalterPaes/go-users-api/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	repository repositories.UserRepository
	jwtAuth    *jwt.Auth
}

func NewLoginHandler(r repositories.UserRepository, jwtAuth *jwt.Auth) *LoginHandler {
	return &LoginHandler{
		repository: r,
		jwtAuth:    jwtAuth,
	}
}

// User Login godoc
// @Summary Login an user
// @Description Login an user
// @Tags login
// @Accept json
// @Produce json
// @Param request body dtos.LoginInput true "Login Request"
// @Success 200 {object} dtos.JWTTokenOutput
// @Failure 400 {array} errors.ValidationError
// @Failure 401
// @Failure 404 {object} errors.CustomError
// @Router /login [post]
func (h *LoginHandler) Login(c *gin.Context) {
	var loginDto dtos.LoginInput

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.NewValidationErrors(err))
		return
	}

	user, err := h.repository.FindByEmail(loginDto.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	if !user.ValidatePassword(loginDto.Password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := h.jwtAuth.GenerateToken(map[string]string{
		"id":       user.ID.String(),
		"username": user.Name,
		"email":    user.Email,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	c.JSON(http.StatusOK, dtos.JWTTokenOutput{Token: token})
}
