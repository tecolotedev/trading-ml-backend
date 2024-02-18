package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/controllers"
	"github.com/tecolotedev/trading-ml-backend/middlewares"
)

func SetUpUserRoutes(router fiber.Router) {
	router.Post("/login", controllers.Login)
	router.Post("/signup", controllers.Signup)
	router.Get("/verify-account", controllers.VerifyAccount)

	// Protected routes
	router.Use(middlewares.Auth)
	router.Put("/user", controllers.UpdateUser)
	router.Delete("/user", controllers.DeleteUser)
}
