package main

import (
	"fmt"

	"github.com/mirjalilova/api_gateway/api"
	"github.com/mirjalilova/api_gateway/api/handlers"
	cf "github.com/mirjalilova/api_gateway/config"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()

	memoryConn, err := grpc.NewClient(fmt.Sprintf("memory%s", config.MEMORY_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Error while connecting to memory service: " + err.Error())
		return
	}
	defer memoryConn.Close()

	timelineConn, err := grpc.NewClient(fmt.Sprintf("timeline%s", config.TIMELINE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Error while connecting to timeline service: " + err.Error())
		return
	}
	defer memoryConn.Close()

	userConn, err := grpc.NewClient(fmt.Sprintf("auth%s", config.USER_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Error while connecting to user service: " + err.Error())
		return
	}
	defer memoryConn.Close()

	h := handlers.NewHandler(memoryConn, timelineConn, userConn)

	r := api.Engine(h)

	slog.Info("Starting HTTP server on port " + config.API_GETEWAY + "...")
	if err := r.Run(config.API_GETEWAY); err != nil {
		slog.Error("Error while starting HTTP server: " + err.Error())
	}
}
