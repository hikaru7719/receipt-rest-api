package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/registry"
)

func main() {
	r := gin.Default()
	container := &registry.Container{}
	v1 := r.Group("/v1")
	{
		v1.GET("/receipt/:id", container.NewAppHandler().GetReceipt())
	}

	r.Run(":8080")

}
