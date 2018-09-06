package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func SlackRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		challenge := make(map[string]interface{})
		context.BindJSON(&challenge)
		fmt.Println(challenge)
		if challenge["token"] == os.Getenv("Slack_Token") {
			context.JSON(200, challenge["challenge"])
		} else {
			context.JSON(500, "error")
		}

	})
	return router
}
