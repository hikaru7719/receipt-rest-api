package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

// SlackHandler - Slack Botのハンドラ
type SlackHandler interface {
	Adapter(ctx *gin.Context)
}

type slackHandler struct {
}

// NewSlackHandler - SlackHandlerの生成
func NewSlackHandler() SlackHandler {
	return &slackHandler{}
}

// Adapter - イベントに応じて振る舞いを決定
func (h *slackHandler) Adapter(context *gin.Context) {
	request := make(map[string]interface{})
	context.BindJSON(&request)
	fmt.Println(request)
	fmt.Println()
	fmt.Println()
	switch request["type"] {
	case "url_verification":
		if request["token"] == os.Getenv("Slack_Token") {
			context.JSON(200, request["challenge"])
		} else {
			context.JSON(500, "error")
		}
	case "event_callback":
		event, _ := request["event"].(map[string]interface{})
		if event["type"] == "app_mention" {
			fmt.Println(event["text"])
		}
	}

}
