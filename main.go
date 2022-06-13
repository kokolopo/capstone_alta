package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kokolopo/capstone_alta/auth"
	"github.com/kokolopo/capstone_alta/config"
	"github.com/kokolopo/capstone_alta/database"
	"github.com/kokolopo/capstone_alta/domain/client"
	"github.com/kokolopo/capstone_alta/domain/user"
	"github.com/kokolopo/capstone_alta/handler"
	"github.com/kokolopo/capstone_alta/routes"
)

func main() {

	conf := config.InitConfiguration()
	database.InitDatabase(conf)
	db := database.DB

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	clientRepo := client.NewClientRepository(db)
	clientService := client.NewUserService(clientRepo)
	clientHandler := handler.NewClientHandler(clientService, userService, authService)

	router := gin.Default()

	router.Use(auth.CORSMiddleware())

	// routes.APIRoutes(router, userHandler, clientHandler, authService, userService)
	routes.APIRoutes(
		router,
		userHandler,
		clientHandler,
		authService,
		userService,
	)

	router.Run()

}
