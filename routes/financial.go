package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/controllers"
)

func SetUpFinancialRoutes(router fiber.Router) {
	router = router.Group("/financial")

	router.Get("/data", controllers.GetFinancialData)
}
