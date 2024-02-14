package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tecolotedev/stori_back/config"
	"github.com/tecolotedev/stori_back/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.SetUpConfig()
	// db.InitDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://stori-front.tecolotedev.com",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ok": true, "message": "api is working"})
	})

	app.Get("/get-token", func(c *fiber.Ctx) error {
		cookie := fiber.Cookie{
			Name:     "access_token",
			Value:    "token_value_get",
			SameSite: "None",
			HTTPOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		}

		c.Cookie(&cookie)
		return c.JSON(fiber.Map{"ok": true, "message": "sending token"})
	})

	app.Post("/post-token", func(c *fiber.Ctx) error {
		cookie := fiber.Cookie{
			Name:     "access_token",
			Value:    "token_value_post",
			SameSite: "None",
			HTTPOnly: true,
			Expires:  time.Now().Add(24 * time.Hour),
		}

		c.Cookie(&cookie)
		return c.JSON(fiber.Map{"ok": true, "message": "sending token"})
	})

	routes.SetUpRoutes(app)

	log.Fatal((app.Listen(":" + config.EnvVars.PORT)))

}
