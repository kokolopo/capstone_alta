package handler

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *ClientHandler) GetClients(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	clients, err := h.clientService.GetAll(currentUser.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Gagal mendapatkan data Clients!", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := client.FormatClients(clients)
	res := helper.ApiResponse("Berhasil mendapatkan data Clients!", http.StatusOK, "berhasil", formatter)

	c.JSON(http.StatusOK, res)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)

	client, err := h.clientService.GetByID(id)
	if err != nil {
		res := helper.ApiResponse("Item Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if client.ID == 0 {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if currentUser.ID != client.UserID {
		res := helper.ApiResponse("Failed to Delete Client", http.StatusBadRequest, "failed", errors.New("kamu bukan tidak berhak"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	_, errDel := h.clientService.DeleteClient(client.ID)
	if errDel != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", errDel)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	cekItem, errCek := h.clientService.GetByID(id)
	if errCek != nil {
		res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", errCek)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if cekItem.ID == 0 {
		res := helper.ApiResponse("Successfuly Delete Item", http.StatusOK, "success", nil)

		c.JSON(http.StatusOK, res)
		return
	}

	data := gin.H{"is_deleted": true}
	res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", data)

	c.JSON(http.StatusCreated, res)
}
