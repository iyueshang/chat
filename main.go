package main

import (
	"github.com/iyueshang/chat/routers"
	"log"
	"net/http"
)

func init()  {
	
}

func main() {
	// Disable log's color
	//gin.DisableConsoleColor()
	//gin.SetMode("debug")

	router := routers.SetupRouter()
	
	server := &http.Server{
		Addr:				":8080",
		Handler:			router,
		ReadHeaderTimeout:	60,
		WriteTimeout:		60,
		MaxHeaderBytes: 	1 << 20,
	}
	log.Fatal(server.ListenAndServe())

}
