package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func Auth(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	payload, err := utils.VerifyToken(accessToken)
	if err != nil {
		c.ClearCookie("access_token")
		fmt.Println("err mid Auth: ", err)
		return err
	}
	c.Locals("userId", payload.USERID)
	return c.Next()
}
