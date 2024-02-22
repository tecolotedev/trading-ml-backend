package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func Auth(c *fiber.Ctx) error {
	// Get token from cookie
	access_token := c.Cookies("access_token")
	if access_token == "" {
		return utils.SendError(c, fmt.Errorf("no token, please signin"), fiber.StatusForbidden)
	}

	// Verify is a valid token
	payload, err := utils.VerifyToken(access_token)
	if err != nil {
		c.ClearCookie("access_token")
		return utils.SendError(c, fmt.Errorf("invalid token, please signin"), fiber.StatusUnauthorized)
	}

	// Send the user id to protected routes
	c.Locals("userID", payload.USERID)

	return c.Next()
}
