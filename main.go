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
	router.Use(CORSMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pang-ping-pong",
		})
	})
	api := router.Group("/api/v1")

	// user
	api.POST("/users", userHandler.UserRegister)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/sessions", userHandler.Login)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	// router.Use(cors.Default())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://invoiceinaja-test.herokuapp.com"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://invoiceinaja-test.herokuapp.com/api/v1"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	//routes.APIRoutes(router, userHandler, authService, userService)

	router.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
