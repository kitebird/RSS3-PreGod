package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/cache"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.Logger()

func init() {
	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup err: %v", err)
	}

	if err := cache.Setup(); err != nil {
		log.Fatalf("cache.Setup err: %v", err)
	}

	if err := db.Setup(); err != nil {
		log.Fatalf("db.Setup err: %v", err)
	}
}

func main() {
	gin.SetMode(config.HubServer.RunMode)

	port := fmt.Sprintf(":%d", config.HubServer.HttpPort)

	server := &http.Server{
		Addr:           port,
		Handler:        router.InitRouter(),
		ReadTimeout:    config.HubServer.ReadTimeout,
		WriteTimeout:   config.HubServer.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	log.Infof("[info] start http server listening %s", port) // TODO: change to zap

	go server.ListenAndServe()

	gracefullyExit(server)
}

func gracefullyExit(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-quit

	log.Infof("Shutdown due to a signal", sig)

	now := time.Now()

	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second) // with a 5s timeout
	defer cancel()

	if err := server.Shutdown(cxt); err != nil {
		log.Fatal("Shutdown error:", err)
	}

	log.Infof("Shutdown server successfully in", time.Since(now))
}
