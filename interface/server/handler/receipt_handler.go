package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
	"strconv"
)

// ReceiptHandler - レシートに関するハンドラ
type ReceiptHandler interface {
	GetReceipt(ctx *gin.Context)
	PostReceipt(ctx *gin.Context)
	DeleteReceipt(ctx *gin.Context)
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
		c.Status(500)
	} else {
		c.JSON(200, receipt)
	}
}

// PostReceipt - レシートの新規作成
func (h *receiptHandler) PostReceipt(c *gin.Context) {
	form := &form.ReceiptForm{}
	err := c.BindJSON(form)
	receipt, err := h.u.PostReceipt(form.UserID, form.Name, form.Price, form.Kind, form.Date, form.Memo, form.CreditID)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(201, receipt)
	}
}

// DeleteReceipt - レシートの削除
func (h *receiptHandler) DeleteReceipt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.u.DeleteReceipt(id)
	if err != nil {
		c.Status(204)
	} else {
		c.Status(202)
	}
}
