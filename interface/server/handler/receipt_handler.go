package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
	"strconv"
)

// ReceiptHandler - レシートに関するハンドラ
type ReceiptHandler interface {
	GetReceipt(ctx *gin.Context)
	PostReceipt(ctx *gin.Context)
}

type receiptHandler struct {
	u usecase.ReceiptUsecase
}

// NewReceiptHandler - receiptHandlerの生成
func NewReceiptHandler(u usecase.ReceiptUsecase) ReceiptHandler {
	return &receiptHandler{u: u}
}

// GetReceipt - レシートの取得
func (h *receiptHandler) GetReceipt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	receipt, err := h.u.GetReceipt(id)
	if err != nil {
		fmt.Println(err)
		c.Status(500)
	} else {
		c.JSON(200, receipt)
	}
}

// PostReceipt - レシートの新規作成
func (h *receiptHandler) PostReceipt(c *gin.Context) {
	form := &form.ReceiptForm{}
	err := c.BindJSON(form)
	receipt, err := h.u.PostReceipt(form.Name, form.Price, form.Kind, form.Date, form.Memo)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, err)
	} else {
		c.JSON(201, receipt)
	}
}
