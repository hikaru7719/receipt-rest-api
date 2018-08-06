package handler

import (
	"github.com/gin-gonic/gin"
)
type GetReceiptHandler interface {
	GetReceipt() gin.HandlerFunc
}

type getReceiptHandler struct{

}

func NewGetReceiptHandler() GetReceiptHandler{
	return &getReceiptHandler{}
}


func (s *getReceiptHandler) GetReceipt() gin.HandlerFunc{
	return func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, "Hello %s", id)
	}
}