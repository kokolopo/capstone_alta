package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokolopo/capstone_alta/auth"
	"github.com/kokolopo/capstone_alta/domain/client"
	"github.com/kokolopo/capstone_alta/domain/user"
	"github.com/kokolopo/capstone_alta/helper"
)

type ClientHandler struct {
	clientService client.IService
	userService   user.IService
	authService   auth.Service
}

func NewClientHandler(ClientService client.IService, userService user.IService, authService auth.Service) *ClientHandler {
	return &ClientHandler{ClientService, userService, authService}
}

func (h *ClientHandler) AddClient(c *gin.Context) {
	//
	var input client.InputAddClient

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login Gagal!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// // didapatkan dari JWT
	currentUser := c.MustGet("currentUser").(user.User)

	_, errClient := h.clientService.AddClient(currentUser.ID, input)
	if errClient != nil {
		errors := helper.FormatValidationError(errClient)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Menambahkan Client Baru Gagal!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"status": "Berhasil Menambahkan Client Baru!",
	}

	res := helper.ApiResponse("Berhasil Membuat client Baru!", http.StatusCreated, "berhasil", data)

	c.JSON(http.StatusCreated, res)
}
