package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokolopo/capstone_alta/auth"
	"github.com/kokolopo/capstone_alta/domain/user"
	"github.com/kokolopo/capstone_alta/handler"
)

func APIRoutes(router *gin.Engine, userHandler *handler.UserHandler, authService auth.Service, userService user.IService) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "wellcome",
		})
	})

	api := router.Group("/api/v1")

	// user
	api.POST("/users", userHandler.UserRegister)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/sessions", userHandler.Login)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.PUT("/users", auth.AuthMiddleware(authService, userService), userHandler.UpdateUser)
	api.POST("/reset_passwords", userHandler.ResetPassword)
}
