package main

import (
	"github.com/hikaru7719/receipt-rest-api/interface/server/router"
)

func main() {

	r := router.Router()
	r.Run(":8080")

}
