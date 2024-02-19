package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/controllers"
)

func SetUpPlansRoutes(router fiber.Router) {

	router.Get("/", controllers.GetPlans)
	router.Get("/:id", controllers.GetPlan)

}
