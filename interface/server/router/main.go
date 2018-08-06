package main

import(
	"github.com/gin-gonic/gin"
	"../../../registry"
)

func main(){
	r := gin.Default()
	con := &registry.Container{}
	v1 := r.Group("/v1"){
		v1.GET("/receipt/:id", con.NewAppHandler().GetReceipt())
	}

	r.Run()

}