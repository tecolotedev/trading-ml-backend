package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tecolotedev/trading-ml-backend/config"
	"github.com/tecolotedev/trading-ml-backend/db"
	"github.com/tecolotedev/trading-ml-backend/email"
	"github.com/tecolotedev/trading-ml-backend/routes"
	"github.com/tecolotedev/trading-ml-backend/utils"

	"github.com/gofiber/fiber/v2"
)

var pack = "main"

func main() {
	// Load env vars
	config.SetUpConfig()

	// Init postgres db
	db.InitDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
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

	// universal wg to block until all go routines have finished
	wg := sync.WaitGroup{}

	email.Mailer.WG = &wg

	// Listen Channels
	go email.Mailer.ListenForEmails()

	go listenForShutdown(&wg)

	utils.Log.FatalLog(app.Listen(":"+config.EnvVars.PORT), pack)

}

func listenForShutdown(wg *sync.WaitGroup) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdown(wg)
	os.Exit(0)
}

func shutdown(wg *sync.WaitGroup) {
	// block until all go-routines have finished
	wg.Wait()

	// break loop of chans
	email.Mailer.MailDoneChan <- true

	// close channels
	close(email.Mailer.MailChan)
	close(email.Mailer.MailDoneChan)

}
