package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"strconv"
)

type GetReceiptHandler interface {
	GetReceipt() gin.HandlerFunc
}

type getReceiptHandler struct {
	u usecase.ReceiptUsecase
}

func NewGetReceiptHandler(u usecase.ReceiptUsecase) GetReceiptHandler {
	return &getReceiptHandler{u}
}

func (h *getReceiptHandler) GetReceipt() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		receipt, _ := h.u.GetReceipt(id)
		c.JSON(200, receipt)
	}
}
