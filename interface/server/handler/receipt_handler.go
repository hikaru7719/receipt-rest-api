package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
	"strconv"
)

type ReceiptHandler interface {
	GetReceipt()
	PostReceipt()
}

type receiptHandler struct {
	u usecase.ReceiptUsecase
}

func NewReceiptHandler(u usecase.ReceiptUsecase) *receiptHandler {
	return &receiptHandler{u}
}

func (h *receiptHandler) GetReceipt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	receipt, _ := h.u.GetReceipt(id)
	c.JSON(200, receipt)
}

func (h *receiptHandler) PostReceipt(c *gin.Context) {
	form := &form.ReceiptForm{}
	err := c.BindJSON(form)
	receipt, err := h.u.PostReceipt(form.Name, form.Kind, form.Date, form.Memo)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(201, receipt)
	}
}
