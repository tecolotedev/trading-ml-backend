package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {
	router := app.Group("/api/user")
	SetUpUserRoutes(router)

	router = app.Group("/api/plan")
	SetUpPlansRoutes(router)

	router = app.Group("/api/financial")
	SetUpFinancialRoutes(router)
}
