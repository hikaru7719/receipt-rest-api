package handler

import (
	"github.com/gin-gonic/gin"
)
type GetReceiptHandler interface {
	GetReceipt(c *gin.Context)
}

type getReceiptHnadler struct{

}

func NewGetReceiptHandler() GetReceiptHandler{
	return &getReceiptHnadler{}
}


func (s *getReceiptHnadler) GetReceipt(c *gin.Context){
	id := c.Params("id")
	c.String(200,"Hello %s",id)
}