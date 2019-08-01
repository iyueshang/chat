package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iyueshang/chat/routers"
	"log"
	"net/http"
)

func init()  {
	
}

func main() {
	// Disable log's color
	gin.DisableConsoleColor()
	gin.SetMode("debug")

	router := routers.SetupRouter()

	server := &http.Server{
		Addr:				fmt.Sprintf(":%d", 8080),
		Handler:			router,
		ReadHeaderTimeout:	60,
		WriteTimeout:		60,
		MaxHeaderBytes: 	1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
