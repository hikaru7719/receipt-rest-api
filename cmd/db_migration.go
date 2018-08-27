package main

import (
	"github.com/hikaru7719/receipt-rest-api/domain/model"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
)

func main() {
	datastore.CreateConnection(datastore.GetDBEnv())
	datastore.DB.AutoMigrate(&model.Credit{})
	datastore.DB.AutoMigrate(&model.Receipt{})
}
