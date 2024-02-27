package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tecolotedev/trading-ml-backend/indicators"
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

func GetSMA(c *fiber.Ctx) error {

	// get params for this indicator
	fillNA := c.Query("fill_na", "Drop")
	seriesType := c.Query("series_type", "close")
	periods := c.QueryInt("periods", 10)

	bars := new([]utils.Bar)

	if err := c.BodyParser(bars); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	if len(*bars) <= periods {
		return utils.SendError(c, fmt.Errorf("bars lenght needs to be greater than periods"), fiber.StatusBadRequest)
	}

	input := indicators.SMAInput{
		SeriesType: seriesType,
		Values:     *bars,
		FillNA:     fillNA,
		Periods:    periods,
	}
	output := indicators.GetSMA(input)

	return utils.SendResponse(c, output)
}

func GetEMA(c *fiber.Ctx) error {

	// get params for this indicator
	fillNA := c.Query("fill_na", "Drop")
	seriesType := c.Query("series_type", "close")
	periods := c.QueryInt("periods", 10)

	bars := new([]utils.Bar)

	if err := c.BodyParser(bars); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	if len(*bars) <= periods {
		return utils.SendError(c, fmt.Errorf("bars lenght needs to be greater than periods"), fiber.StatusBadRequest)
	}

	input := indicators.EMAInput{
		SeriesType: seriesType,
		Values:     *bars,
		FillNA:     fillNA,
		Periods:    periods,
	}
	output := indicators.GetEMA(input)

	return utils.SendResponse(c, output)
}

func GetRSI(c *fiber.Ctx) error {

	// get params for this indicator
	fillNA := c.Query("fill_na", "Drop")
	seriesType := c.Query("series_type", "close")
	periods := c.QueryInt("periods", 14)

	bars := new([]utils.Bar)

	if err := c.BodyParser(bars); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	if len(*bars) <= periods {
		return utils.SendError(c, fmt.Errorf("bars lenght needs to be greater than periods"), fiber.StatusBadRequest)
	}

	input := indicators.RSIInput{
		SeriesType: seriesType,
		Values:     *bars,
		FillNA:     fillNA,
		Periods:    periods,
	}
	output := indicators.GetRSI(input)

	return utils.SendResponse(c, output)
}

func GetMACD(c *fiber.Ctx) error {

	// get params for this indicator
	fillNA := c.Query("fill_na", "Drop")
	seriesType := c.Query("series_type", "close")
	fastPeriods := c.QueryInt("fast_periods", 12)
	slowPeriods := c.QueryInt("slow_periods", 26)
	macdPeriods := c.QueryInt("macd_periods", 9)

	bars := new([]utils.Bar)

	if err := c.BodyParser(bars); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	if len(*bars) <= fastPeriods && len(*bars) <= slowPeriods && len(*bars) <= macdPeriods {
		return utils.SendError(c, fmt.Errorf("bars lenght needs to be greater than periods"), fiber.StatusBadRequest)
	}

	input := indicators.MACDInput{
		SeriesType:  seriesType,
		Values:      *bars,
		FillNA:      fillNA,
		SlowPeriods: slowPeriods,
		FastPeriods: fastPeriods,
		MACDPeriods: macdPeriods,
	}
	output := indicators.GetMACD(input)

	return utils.SendResponse(c, output)
}

func GetBBANDS(c *fiber.Ctx) error {
	// get params for this indicator
	fillNA := c.Query("fill_na", "Drop")
	seriesType := c.Query("series_type", "close")
	periods := c.QueryInt("periods", 20)
	sd := c.QueryInt("sd", 2)

	bars := new([]utils.Bar)

	if err := c.BodyParser(bars); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	if len(*bars) <= periods {
		return utils.SendError(c, fmt.Errorf("bars lenght needs to be greater than periods"), fiber.StatusBadRequest)
	}

	input := indicators.BBANDSInput{
		SeriesType: seriesType,
		Values:     *bars,
		FillNA:     fillNA,
		Periods:    periods,
		SD:         sd,
	}
	output := indicators.GetBBANDS(input)

	return utils.SendResponse(c, output)
}

func GetPTC(c *fiber.Ctx) error {
	// get params for this indicator
	fillNA := c.Query("fill_na", "Drop")
	seriesType := c.Query("series_type", "close")

	bars := new([]utils.Bar)

	if err := c.BodyParser(bars); err != nil {
		utils.Log.ErrorLog(err, pack)
		return utils.SendError(c, fmt.Errorf("error in body request"), fiber.StatusBadRequest)
	}

	input := indicators.PTCInput{
		Values:     *bars,
		FillNA:     fillNA,
		SeriesType: seriesType,
	}
	output := indicators.GetPTC(input)

	return utils.SendResponse(c, output)
}
