package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"server-http/config"
	"server-http/logger"
	"server-http/server"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
		return
	}
	logger.InitLogger()
	defer zap.L().Sync() // Ensure the logs are flushed when the application ends

	r := server.SetupRoutes()

	port := strconv.Itoa(cfg.GetPort())
	appServer := server.NewHTTPServer(":"+port, r)

	if err := appServer.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// Listen for OS signals to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait for a signal

	if err := appServer.Stop(); err != nil {
		fmt.Println("Error stopping server:", err)
	}
}
