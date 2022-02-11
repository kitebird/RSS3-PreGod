package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/routers"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/cache"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/db"
	"github.com/gin-gonic/gin"
)

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
		Handler:        routers.InitRouter(),
		ReadTimeout:    config.HubServer.ReadTimeout,
		WriteTimeout:   config.HubServer.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	log.Printf("[info] start http server listening %s", port) // TODO: change to zap

	go server.ListenAndServe()

	gracefullyExit(server)
}

func gracefullyExit(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-quit

	log.Println("Shutdown due to a signal", sig)

	now := time.Now()

	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second) // with a 5s timeout
	defer cancel()

	if err := server.Shutdown(cxt); err != nil {
		log.Fatal("Shutdown error:", err)
	}

	log.Println("Shutdown server successfully in", time.Since(now))
}
