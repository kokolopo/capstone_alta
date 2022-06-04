package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokolopo/capstone_alta/auth"
	"github.com/kokolopo/capstone_alta/config"
	"github.com/kokolopo/capstone_alta/database"
	"github.com/kokolopo/capstone_alta/domain/user"
	"github.com/kokolopo/capstone_alta/handler"
)

func main() {

	conf := config.InitConfiguration()
	database.InitDatabase(conf)
	db := database.DB

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := router.Group("/api/v1")

	// user
	api.POST("/users", userHandler.UserRegister)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/sessions", userHandler.Login)

	router.Run()

}
