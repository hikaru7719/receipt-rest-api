package main

import(
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	v1 := r.Group("/v1"){
		v1.GET("/receipt/:id",)
	}

	r.Run()

}