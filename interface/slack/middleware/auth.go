package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Auth(ctx *gin.Context) {
	request := make(map[string]interface{})
	ctx.BindJSON(&request)
	if token, _ := request["token"].(string); token == os.Getenv("Slack_Token") {
		ctx.Set("request", request)
		ctx.Next()
	} else {
		ctx.JSON(400, "Bad Request")
	}
}
