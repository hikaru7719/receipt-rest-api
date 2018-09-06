package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/form"
	"os"
)

func SlackRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		challenge := form.HttpChallenge{}
		context.BindJSON(&challenge)
		fmt.Println(challenge)
		if challenge.Token == os.Getenv("Slack_Token") {
			context.JSON(200, challenge.Challenge)
		} else {
			context.JSON(500, "error")
		}

	})
	return router
}
