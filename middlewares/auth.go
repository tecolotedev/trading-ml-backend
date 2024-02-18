package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func Auth(c *fiber.Ctx) error {
	access_token := c.Cookies("access_token")
	if access_token == "" {
		return utils.SendError(c, fmt.Errorf("no token, please signin"), fiber.StatusForbidden)
	}

	payload, err := utils.VerifyToken(access_token)
	if err != nil {
		c.ClearCookie("access_token")
		return utils.SendError(c, fmt.Errorf("invalid token, please signin"), fiber.StatusUnauthorized)
	}

	c.Locals("userID", payload.USERID)

	return c.Next()
}
