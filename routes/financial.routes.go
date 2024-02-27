package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/controllers"
	"github.com/tecolotedev/trading-ml-backend/middlewares"
)

func SetUpFinancialRoutes(router fiber.Router) {
	router.Use(middlewares.Auth, middlewares.FinancialMiddleware)
	router.Get("/data", controllers.GetFinancialData)
	router.Get("/indicator/sma", controllers.GetSMA)
	router.Get("/indicator/ema", controllers.GetEMA)
	router.Get("/indicator/rsi", controllers.GetRSI)

	// router.Get("/indicator/macd", controllers.GetMACDData)
	// router.Get("/indicator/bbands", controllers.GetBBANDSData)
}
