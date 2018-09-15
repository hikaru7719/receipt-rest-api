package handler

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/parser"
	"os"
	"strings"
)

// SlackHandler - Slack Botのハンドラ
type SlackHandler interface {
	Adapter(ctx *gin.Context)
}

type slackHandler struct {
	parser parser.Parser
}

// NewSlackHandler - SlackHandlerの生成
func NewSlackHandler(parser parser.Parser) SlackHandler {
	return &slackHandler{parser: parser}
}

// Adapter - イベントに応じて振る舞いを決定
func (h *slackHandler) Adapter(context *gin.Context) {
	request, _ := context.Get("request")
	requestJson, _ := request.(map[string]interface{})
	requestJsonString, _ := requestJson["type"].(string)
	switch requestJsonString {
	case "url_verification":
		if requestJson["token"] == os.Getenv("Slack_Token") {
			context.JSON(200, requestJson["challenge"])
		} else {
			context.JSON(500, "error")
		}
	case "event_callback":
		event, _ := requestJson["event"].(map[string]interface{})
		eventString, _ := event["type"].(string)
		if eventString == "app_mention" {
			text, _ := event["text"].(string)
			stringReader := strings.NewReader(text)
			scanner := bufio.NewScanner(stringReader)
			scanner.Scan()
			scanner.Scan()
			switch text := scanner.Text(); text {
			case "クレジット登録":
				fmt.Println(text)
			}
		}
	}
}
