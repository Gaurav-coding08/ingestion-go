package main

import (
	"github.com/Gaurav-coding08/ingestion-go/cmd/kafka"
	"github.com/Gaurav-coding08/ingestion-go/cmd/server"
	"github.com/Gaurav-coding08/ingestion-go/config"
	"github.com/Gaurav-coding08/ingestion-go/database/connect"
)

func main() {
	cfg := config.LoadConfig()

	connect.InitDB(cfg)
	db := connect.DB

	// Init Kafka
	producer, err := kafka.ConnectToProducer(cfg)
	if err == nil {
		server.StartServer(cfg, producer, db)
	}
}
