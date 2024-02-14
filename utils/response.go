package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func SendError(c *fiber.Ctx, err error, status int) error {
	c.Status(status)
	res := ErrorResponse{Ok: false, Message: err.Error()}
	return c.JSON(res)
}

type Response struct {
	Ok   bool `json:"ok"`
	Data any  `json:"data"`
}

func SendResponse(c *fiber.Ctx, data any) error {
	res := Response{Ok: true, Data: data}
	return c.JSON(res)
}
