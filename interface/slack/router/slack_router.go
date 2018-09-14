package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/handler"
	"github.com/hikaru7719/receipt-rest-api/interface/slack/middleware"
)

func SlackRouter() *gin.Engine {
	slackHandler := handler.NewSlackHandler()
	router := gin.Default()
	router.Use(middleware.Auth)
	router.POST("/", slackHandler.Adapter)
	return router
}
