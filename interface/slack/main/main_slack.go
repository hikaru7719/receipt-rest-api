package main

import (
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/router"
)

func main() {
	connect := datastore.GetDBEnv()
	r := router.SlackRouter()
	datastore.CreateConnection(connect)
	r.Run(":8080")
}
