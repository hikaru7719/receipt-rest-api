package main

import (
	"fmt"
	"github.com/hikaru7719/receipt-rest-api/infrastructure/datastore"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/router"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MYSQL_IP"))
	connect := datastore.GetDBEnv()
	fmt.Println(connect)
	r := router.SlackRouter()
	datastore.CreateConnection(connect)
	r.Run(":8080")

}
