package dtos

import "github.com/WalterPaes/go-users-api/internal/entity"

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserInput struct {
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

type UsersListOutput struct {
	Users   []entity.User `json:"users"`
	Page    int           `json:"page"`
	PerPage int           `json:"per_page"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
