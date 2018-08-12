package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"strconv"
	"fmt"
)

type ReceiptHandler interface {
	GetReceipt() gin.HandlerFunc
	PostReceipt() gin.HandlerFunc
}

type receiptHandler struct {
	u usecase.ReceiptUsecase
}

func NewReceiptHandler(u usecase.ReceiptUsecase) *receiptHandler {
	return &receiptHandler{u}
}

func (h *receiptHandler) GetReceipt() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		receipt, _ := h.u.GetReceipt(id)
		c.JSON(200, receipt)
	}
}

func (h *receiptHandler) PostReceipt() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print(c.Request.Body)
		name := c.PostForm("name")
		kind := c.PostForm("kind")
		date := c.PostForm("date")
		memo := c.PostForm("memo")
		receipt,err := h.u.PostReceipt(name, kind, date, memo)

		fmt.Println("abc",name,kind,date,memo)
		if err != nil {
			c.JSON(500,err)
		}
		c.JSON(201, receipt)
	}
}
