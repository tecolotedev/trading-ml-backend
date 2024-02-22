package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/controllers"
	"github.com/tecolotedev/trading-ml-backend/middlewares"
)

func SetUpFinancialRoutes(router fiber.Router) {
	router.Use(middlewares.Auth, middlewares.FinancialMiddleware)
	router.Get("/data", controllers.GetFinancialData)
	router.Get("/indicator/ma", controllers.GetMAData)

}
