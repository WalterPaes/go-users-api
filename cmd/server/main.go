package main

import (
	"github.com/WalterPaes/go-users-api/internal/entity"
	"github.com/WalterPaes/go-users-api/internal/handlers"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entity.User{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userHandler := handlers.NewUserHandler(repositories.NewUserRepository(db))

	r.GET("/users", userHandler.FindAll)
	r.GET("/users/:id", userHandler.FindUserById)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.Delete)

	r.Run(":8001")
}
