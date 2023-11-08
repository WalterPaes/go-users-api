package main

import (
	"github.com/WalterPaes/go-users-api/internal/entity"
	"github.com/WalterPaes/go-users-api/internal/handlers"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	"github.com/WalterPaes/go-users-api/pkg/middlewares"
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

	loginHandler := handlers.NewLoginHandler(repositories.NewUserRepository(db))
	userHandler := handlers.NewUserHandler(repositories.NewUserRepository(db))

	users := r.Group("/users")
	users.Use(middlewares.JwtAuthMiddleware())
	users.GET("/", userHandler.FindAll)
	users.GET("/:id", userHandler.FindUserById)
	users.POST("/", userHandler.CreateUser)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.Delete)

	r.POST("/login", loginHandler.Login)

	r.Run(":8001")
}
