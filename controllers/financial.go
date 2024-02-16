package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/twelve_data"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func GetFinancialData(c *fiber.Ctx) error {

	// Get and validate symbol
	symbol := c.Query("symbol")
	if isValid := utils.ValidateSymbol(symbol); !isValid {
		return utils.SendError(c, fmt.Errorf("invalid symbol"), fiber.StatusBadRequest)
	}

	// Get and validate interval
	interval := c.Query("interval")
	if isValid := utils.ValidateInterval(interval); !isValid {
		return utils.SendError(c, fmt.Errorf("invalid interval"), fiber.StatusBadRequest)
	}

	// Get and validate outputsize
	outputSize := c.QueryInt("outputsize", 30)
	if isValid := utils.ValidateOutputSize(outputSize); !isValid {
		return utils.SendError(c, fmt.Errorf("invalid outputsize"), fiber.StatusBadRequest)
	}

	// Get and validate time zone
	tz := c.Query("time_zone", "UTC")
	_, err := time.LoadLocation(tz)
	if err != nil {
		return utils.SendError(c, fmt.Errorf("invalid Time Zone"), fiber.StatusBadRequest)
	}

	// fetch data from twelve data api
	res, err := twelve_data.FetchTimeSeries(outputSize, symbol, interval, tz)

	if err != nil {
		return utils.SendError(c, err, fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, res)
}
