package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"github.com/hikaru7719/receipt-rest-api/interface/server/form"
	"strconv"
)

// CreditHandler - クレジットに関するハンドラ
type CreditHandler interface {
	PostCredit(ctx *gin.Context)
	GetCreditList(ctx *gin.Context)
}

type creditHandler struct {
	u usecase.CreditUsecase
}

// NewCreditHandler - CreditHandlerの生成
func NewCreditHandler(u usecase.CreditUsecase) CreditHandler {
	return &creditHandler{u: u}
}

// PostCredit - クレジットカードの新規作成
func (h *creditHandler) PostCredit(c *gin.Context) {
	form := &form.CreditForm{}
	err := c.BindJSON(form)
	credit, err := h.u.PostCredit(form.UserID, form.CardName, form.FinishDate, form.WithdrawalDate, form.LaterPaymentMonth)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(201, credit)
	}
}

func (h *creditHandler) GetCreditList(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("userID"))
	creditList, err := h.u.GetCreditList(userID)
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, creditList)
	}
}
