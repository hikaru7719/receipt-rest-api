package main

import (
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/server/router"
)

func main() {

	r := router.SlackRouter()
	datastore.CreateConnection(datastore.GetDBEnv())
	r.Run(":8080")

}
