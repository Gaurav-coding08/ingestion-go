package server

import (
	"github.com/Gaurav-coding08/ingestion-go/cmd/kafka"
	"github.com/Gaurav-coding08/ingestion-go/config"
	stocksCntrl "github.com/Gaurav-coding08/ingestion-go/internal/app/controllers/stocks"
	stocksSvc "github.com/Gaurav-coding08/ingestion-go/internal/app/services/stocks"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	authCntrl "github.com/Gaurav-coding08/ingestion-go/internal/app/controllers/auth"
	authRepo "github.com/Gaurav-coding08/ingestion-go/internal/app/repositories/auth"
	authSvc "github.com/Gaurav-coding08/ingestion-go/internal/app/services/auth"
	customMiddleware "github.com/Gaurav-coding08/ingestion-go/pkg/client"
)

func SetupRoutes(r *gin.Engine, cfg *config.AppConfig, producer kafka.KafkaProducer, db *gorm.DB) {

	deps := InitializeDependencies(producer, db)

	addStocksRoutes(r, deps)
	addAuthRoutes(r, deps)
}

type Dependencies struct {
	StocksController *stocksCntrl.Controller
	AuthController   *authCntrl.Controller
}

func InitializeDependencies(producer kafka.KafkaProducer, db *gorm.DB) *Dependencies {
	return &Dependencies{
		StocksController: stocksCntrl.New(
			stocksSvc.New(
				producer,
			),
		),
		AuthController: authCntrl.New(
			authSvc.New(
				authRepo.New(db),
			),
		),
	}
}

func addStocksRoutes(r *gin.Engine, deps *Dependencies) {
	api := r.Group("/api/v1")

	apiSendMessage := api.Group("/send-message")
	{ // only authenticated users which are registered can update
		// the stock price and give live updates to all the clients and consumers
		apiSendMessage.POST("/stocks", customMiddleware.AuthMiddleware(), deps.StocksController.Update)
	}
}

func addAuthRoutes(r *gin.Engine, deps *Dependencies) {

	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", deps.AuthController.Register)
		api.POST("/login", deps.AuthController.Login)
	}

}
