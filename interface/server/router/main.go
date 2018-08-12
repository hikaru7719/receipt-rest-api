package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/server/handler"
)

func main() {

	db := datastore.GetConnection()
	defer db.Close()
	receipt_repository := datastore.NewReceiptRepository(db)
	receipt_usecase := usecase.NewReceiptUsecase(receipt_repository)
	receipt_handller := handler.NewReceiptHandler(receipt_usecase)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/receipt/:id", receipt_handller.GetReceipt())
		v1.POST("/receipt", receipt_handller.PostReceipt())
	}

	r.Run(":8080")

}
