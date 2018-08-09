package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/interface/server/handler"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/receipt/:id", handler.NewGetReceiptHandler().GetReceipt())
	}

	r.Run(":8080")

}
