package server

import (
	"github.com/Gaurav-coding08/ingestion-go/config"
	"github.com/Gaurav-coding08/ingestion-go/cmd/kafka"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(cfg *config.AppConfig, producer kafka.KafkaProducer, db *gorm.DB) {
	r := gin.Default()

	SetupRoutes(r, cfg, producer, db)

	r.Run(":" + cfg.Port)
}
