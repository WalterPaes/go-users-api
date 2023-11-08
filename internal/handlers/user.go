package handlers

import (
	"net/http"

	"github.com/WalterPaes/go-users-api/internal/dtos"
	"github.com/WalterPaes/go-users-api/internal/entity"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	entityId "github.com/WalterPaes/go-users-api/pkg/entity"
	customErrors "github.com/WalterPaes/go-users-api/pkg/errors"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repository repositories.UserRepository
}

func NewUserHandler(r repositories.UserRepository) *UserHandler {
	return &UserHandler{
		repository: r,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userDto dtos.CreateUserInput

	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.NewValidationErrors(err))
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = h.repository.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) FindUserById(c *gin.Context) {
	id := c.Param("id")

	uuid, err := entityId.ParseID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.repository.FindById(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	uuid, err := entityId.ParseID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var userDto dtos.UpdateUserInput
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, customErrors.NewValidationErrors(err))
		return
	}

	user := &entity.User{
		ID:       uuid,
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = h.repository.Update(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	uuid, err := entityId.ParseID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = h.repository.Delete(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
