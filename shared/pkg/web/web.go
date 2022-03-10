package web

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Config struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Handler      http.Handler
}

// Setup starts the web server.
// Returns the address of the server.
func Setup(cfg *Config) string {
	gin.SetMode(cfg.RunMode)

	addr := net.JoinHostPort("localhost", strconv.Itoa(cfg.HttpPort))

	server := &http.Server{
		Addr:           addr,
		Handler:        cfg.Handler,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	go server.ListenAndServe()

	gracefullyExit(server)

	return addr
}

func gracefullyExit(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-quit

	logger.Info("Shutdown due to a signal: ", sig)

	now := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // with a 5s timeout
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Shutdown error:", err)
	}

	logger.Info("Shutdown server successfully in ", time.Since(now))
}
