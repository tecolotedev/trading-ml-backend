package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/stori_back/controllers"
	"github.com/tecolotedev/stori_back/middlewares"
)

func SetUpAccountRoutes(router fiber.Router) {
	router.Use(middlewares.Auth)
	router.Get("/account", controllers.ListAccounts)
	router.Post("/account", controllers.CreateAccount)
	router.Get("/account/:account_id", controllers.GetAccount)
	router.Put("/account/:account_id", controllers.UpdateBalanceAccount)
}
