package main

import (
	"log"

	"github.com/WalterPaes/go-users-api/config"
	"github.com/WalterPaes/go-users-api/internal/entity"
	"github.com/WalterPaes/go-users-api/internal/handlers"
	"github.com/WalterPaes/go-users-api/internal/repositories"
	"github.com/WalterPaes/go-users-api/pkg/jwt"
	"github.com/WalterPaes/go-users-api/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(sqlite.Open(cfg.DbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	jwtAuth := jwt.NewAuth(cfg.JwtSecret, cfg.JwtExpires)

	loginHandler := handlers.NewLoginHandler(repositories.NewUserRepository(db), jwtAuth)
	userHandler := handlers.NewUserHandler(repositories.NewUserRepository(db))

	r := gin.Default()
	r.POST("/login", loginHandler.Login)

	users := r.Group("/users")
	users.Use(middlewares.JwtAuthMiddleware(jwtAuth))
	users.GET("/", userHandler.FindAll)
	users.GET("/:id", userHandler.FindUserById)
	users.POST("/", userHandler.CreateUser)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.Delete)

	r.Run(cfg.AppPort)
}
