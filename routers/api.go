package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iyueshang/chat/ws"
)

var wsController = new(ws.WsController)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	gin.SetMode("debug")

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/ws", func(c *gin.Context) {
		wsController.WsHandler(c.Writer, c.Request)
	})
	router.GET("/wstest", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return router
}