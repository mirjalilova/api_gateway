package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/mirjalilova/api_gateway/api"
	"github.com/mirjalilova/api_gateway/api/handlers"
	cf "github.com/mirjalilova/api_gateway/config"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()

	memoryConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.MEMORY_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Error while connecting to reservation service: " + err.Error())
		return
	}
	defer memoryConn.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping().Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	h := handlers.NewHandler(memoryConn, rdb)

	r := api.Engine(h)

	slog.Info("Starting HTTP server on port " + config.API_GETEWAY + "...")
	if err := r.Run(config.API_GETEWAY); err != nil {
		slog.Error("Error while starting HTTP server: " + err.Error())
	}
}
