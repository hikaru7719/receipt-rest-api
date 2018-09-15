package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/handler"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/middleware"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/parser"
)

func SlackRouter() *gin.Engine {
	slackHandler := handler.NewSlackHandler(parser.NewParser())
	router := gin.Default()
	router.Use(middleware.Auth)
	router.POST("/", slackHandler.Adapter)
	return router
}
