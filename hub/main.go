package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/modules/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/modules/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Setup()
}

func main() {
	gin.SetMode(config.Server.RunMode)

	port := fmt.Sprintf(":%d", config.Server.HttpPort)

	server := &http.Server{
		Addr:           port,
		Handler:        routers.InitRouter(),
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	log.Printf("[info] start http server listening %s", port) // TODO: change to zap

	server.ListenAndServe()
}
