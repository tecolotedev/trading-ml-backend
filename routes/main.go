package routes

import "github.com/gofiber/fiber/v2"

func SetUpRoutes(app *fiber.App) {
	router := app.Group("/api")
	SetUpUserRoutes(router)
	SetUpAccountRoutes(router)
}
