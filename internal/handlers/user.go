package handlers

import (
	"errors"
	"net/http"
	"strconv"

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

// List Users godoc
// @Summary list all users
// @Description list all users
// @Tags users
// @Produce json
// @Param page query string false "page number"
// @Param perPage query string false "items per page number"
// @Success 200 {object} dtos.UsersListOutput
// @Failure 400 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /users [get]
// @Security ApiKeyAuth
func (h *UserHandler) FindAll(c *gin.Context) {
	var err error
	page := c.Param("page")
	perPage := c.Param("perPage")
	pageInt := 1
	perPageInt := 10

	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusBadRequest, customErrors.New("Query Param Error", errors.New("'page' must be a valid int number")))
			return
		}
	}

	if perPage != "" {
		perPageInt, err = strconv.Atoi(perPage)
		if err != nil {
			c.JSON(http.StatusBadRequest, customErrors.New("Query Param Error", errors.New("'per_page' must be a valid int number")))
			return
		}
	}

	users, err := h.repository.FindAll(pageInt, perPageInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, dtos.UsersListOutput{
		Users:   users,
		Page:    pageInt,
		PerPage: perPageInt,
	})
}

// Get User godoc
// @Summary get an user
// @Description get an user
// @Tags users
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} entity.User
// @Failure 400 {array} errors.CustomError
// @Failure 404
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func (h *UserHandler) FindUserById(c *gin.Context) {
	id := c.Param("id")

	uuid, err := entityId.ParseID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.repository.FindById(uuid)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Create User godoc
// @Summary create an user
// @Description create an user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dtos.CreateUserInput true "user request"
// @Success 201 {object} entity.User
// @Failure 400 {array} errors.ValidationError
// @Failure 500 {object} errors.CustomError
// @Router /users [post]
// @Security ApiKeyAuth
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

// Update User godoc
// @Summary update an user
// @Description update an user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param request body dtos.UpdateUserInput true "Update user request"
// @Success 200 {object} entity.User
// @Failure 400 {array} errors.ValidationError
// @Router /users/{id} [put]
// @Security ApiKeyAuth
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

// Delete User godoc
// @Summary delete an user
// @Description delete an user
// @Tags users
// @Produce json
// @Param id path string true "user id"
// @Success 204
// @Failure 500 {object} errors.CustomError
// @Router /users/{id} [delete]
// @Security ApiKeyAuth
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

	c.JSON(http.StatusNoContent, gin.H{})
}
