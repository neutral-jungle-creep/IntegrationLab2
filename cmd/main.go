package main

import (
	v1 "IntegrationLab2/internal/delivery/http/v1"
	"IntegrationLab2/pkg/httpServer"
	"IntegrationLab2/pkg/logger"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := logger.NewLogger()
	log.Info("logger has been started")

	handler := v1.NewHandler(log)

	server := httpServer.NewServer(handler.InitRoutes())

	log.Infof("server started on: [http://localhost:%s]", "8008")

	if err := server.Run(); err != nil {
		log.Fatalf("error http server, %s", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("error on server shutting down, %s", err.Error())
	}
}
