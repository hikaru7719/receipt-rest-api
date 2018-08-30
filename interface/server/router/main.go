package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/application/usecase"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/server/handler"
)

func main() {

	r := router()
	datastore.CreateConnection(datastore.GetDBEnv())
	r.Run(":8080")

}

func router() *gin.Engine {

	receiptRepository := datastore.NewReceiptRepository()
	receiptUseCase := usecase.NewReceiptUsecase(receiptRepository)
	receiptHandler := handler.NewReceiptHandler(receiptUseCase)
	creditRepository := datastore.NewCreditRepository()
	creditUseCase := usecase.NewCreditUsecase(creditRepository)
	creditHandler := handler.NewCreditHandler(creditUseCase)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/receipt/:id", receiptHandler.GetReceipt)
		v1.POST("/receipt", receiptHandler.PostReceipt)
		v1.DELETE("/receipt/:id", receiptHandler.DeleteReceipt)

		v1.POST("/credit", creditHandler.PostCredit)
	}

	return r
}
