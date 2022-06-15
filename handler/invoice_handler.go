package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kokolopo/capstone_alta/auth"
	"github.com/kokolopo/capstone_alta/domain/client"
	"github.com/kokolopo/capstone_alta/domain/invoice"
	"github.com/kokolopo/capstone_alta/domain/user"
	"github.com/kokolopo/capstone_alta/helper"
)

type InvoiceHandler struct {
	invoiceService invoice.IService
	clientService  client.IService
	authService    auth.Service
}

func NewInvoiceHandler(invoiceService invoice.IService, clientService client.IService, authService auth.Service) *InvoiceHandler {
	return &InvoiceHandler{invoiceService, clientService, authService}
}

func (h *InvoiceHandler) AddInvoice(c *gin.Context) {
	// didapatkan dari JWT
	currentUser := c.MustGet("currentUser").(user.User)

	var input invoice.InputAddInvoice
	// tangkap input body
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New Data Has Been Failed1", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	// cek id client relate dengan user'
	client, errClient := h.clientService.GetByID(input.Invoice.ClientID)
	if errClient != nil {
		res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", errClient)

		c.JSON(http.StatusBadRequest, res)
		return
	}
	if client.UserID == 0 {
		res := helper.ApiResponse("client g ada!", http.StatusUnprocessableEntity, "failed", errClient)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	if client.UserID != currentUser.ID {
		res := helper.ApiResponse("ini bukan client anda!", http.StatusUnprocessableEntity, "failed", errClient)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	} else {
		// record data invoice
		newInvoice, errOrder := h.invoiceService.AddInvoice(input)
		if errOrder != nil {
			res := helper.ApiResponse("Invoice Has Been Failed", http.StatusBadRequest, "failed", errOrder)

			c.JSON(http.StatusBadRequest, res)
		}

		// record data detail order
		_, errDetails := h.invoiceService.SaveDetail(newInvoice.ID, input)
		if errDetails != nil {
			res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errDetails)

			c.JSON(http.StatusBadRequest, res)
		}

		data := gin.H{"is_recorded": true}
		res := helper.ApiResponse("Order Has Been Created", http.StatusCreated, "success", data)

		c.JSON(http.StatusCreated, res)
	}

}
