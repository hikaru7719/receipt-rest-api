package main

import (
	"github.com/hikaru7719/receipt-rest-api/interface/slack/router"
)

func main() {
	r := router.SlackRouter()
	r.Run(":8080")
}
