package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/twelve_data"
	"github.com/tecolotedev/trading-ml-backend/utils"
)

func GetFinancialData(c *fiber.Ctx) error {

	// Get and validate symbol
	symbol := c.Query("symbol")

	// Get and validate interval
	interval := c.Query("interval")

	// Get and validate outputsize
	outputSize := c.QueryInt("outputsize", 30)

	// Get start and end date
	startDate := c.Query("start_date", "2006-01-01 00:00:00")
	endDate := c.Query("end_date", time.Now().Format("2006-01-02 15:04:05"))

	// Get and validate time zone
	tz := c.Query("time_zone", "UTC")

	err := utils.ValidateBaseParams(outputSize, symbol, interval, tz, startDate, endDate)
	if err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	// fetch data from twelve data api
	res, err := twelve_data.FetchTimeSeries(outputSize, symbol, interval, tz, startDate, endDate)

	if err != nil {
		return utils.SendError(c, err, fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, res)
}

func GetMAData(c *fiber.Ctx) error {

	// Get and validate symbol
	symbol := c.Query("symbol")

	// Get and validate interval
	interval := c.Query("interval")

	// Get and validate outputsize
	outputSize := c.QueryInt("outputsize", 30)

	// Get start and end date
	startDate := c.Query("start_date", "2006-01-01 00:00:00")
	endDate := c.Query("end_date", time.Now().Format("2006-01-02 15:04:05"))

	// Get and validate time zone
	tz := c.Query("time_zone", "UTC")

	err := utils.ValidateBaseParams(outputSize, symbol, interval, tz, startDate, endDate)
	if err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}

	// get params for this indicator
	maType := c.Query("ma_type", "SMA")
	seriesType := c.Query("series_type", "close")
	timePeriod := c.QueryInt("time_period", 10)

	// validate params for this custom indicator
	err = utils.ValidateMAParams(timePeriod, maType, seriesType)
	if err != nil {
		return utils.SendError(c, err, fiber.StatusBadRequest)
	}
	// fetch data from twelve data api
	res, err := twelve_data.FetchMA(outputSize, timePeriod, symbol, interval, tz, startDate, endDate, maType, seriesType)

	if err != nil {
		return utils.SendError(c, err, fiber.StatusInternalServerError)
	}

	return utils.SendResponse(c, res)
}
