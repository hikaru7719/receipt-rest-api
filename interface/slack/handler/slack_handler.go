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
	request, _ := context.Get("request")
	requestJson, _ := request.(map[string]interface{})
	switch requestJson["type"] {
	case "url_verification":
		if requestJson["token"] == os.Getenv("Slack_Token") {
			context.JSON(200, requestJson["challenge"])
		} else {
			context.JSON(500, "error")
		}
	case "event_callback":
		event, _ := requestJson["event"].(map[string]interface{})
		if event["type"] == "app_mention" {
			fmt.Println(event["text"])
		}
	}
}
