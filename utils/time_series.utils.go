package utils

import (
	"fmt"
	"strconv"
	"time"
)

// Helpers for general financial data
var Symbols = map[string]string{
	"EURUSD": "EUR/USD",
	"GBPUSD": "GBP/USD",
	"AAPL":   "AAPL",
}
var Intervals = map[string]string{
	"1m":  "1min",
	"5m":  "5min",
	"15m": "15min",
	"30m": "30min",
	"45m": "45min",
	"1h":  "1h",
	"2h":  "2h",
	"4h":  "4h",
	"1d":  "1day",
	"1w":  "1week",
	"1M":  "1month",
}

var pack = "utils"

func ValidateSymbol(symbol string) (err error) {
	for key := range Symbols {
		if key == symbol {
			return
		}
	}

	return fmt.Errorf("symbol is not a valid value")
}

func ValidateInterval(interval string) (err error) {
	for key := range Intervals {
		if key == interval {
			return
		}
	}

	return fmt.Errorf("interval is not a valid value")
}
func ValidateDates(startDate, endDate string) (err error) {
	sd, err := time.Parse("2006-01-02 15:04:05", startDate)
	if err != nil {
		Log.ErrorLog(err, pack)
		err = fmt.Errorf("start_date in wrong format")
		return
	}

	ed, err := time.Parse("2006-01-02 15:04:05", endDate)
	if err != nil {
		Log.ErrorLog(err, pack)
		err = fmt.Errorf("end_date in wrong format")
		return
	}

	if sd.After(ed) {
		err = fmt.Errorf("start date is after end date")
		Log.ErrorLog(err, pack)
		return
	}
	if ed.After(time.Now()) {
		err = fmt.Errorf("end_date is after current time")
		Log.ErrorLog(err, pack)
		return
	}

	return

}

func ValidateTimeZone(tz string) (err error) {
	_, err = time.LoadLocation(tz)
	if err != nil {
		Log.ErrorLog(err, pack)
		err = fmt.Errorf("time_zone is invalid")
	}
	return
}

func ValidateOutputSize(outputSize int) (err error) {
	if outputSize > 0 && outputSize <= 5000 {
		return
	}
	return fmt.Errorf("outputsize is invalid")
}

func ValidateBaseParams(outputSize int, symbol, interval, tz, startDate, endDate string) (err error) {
	err = ValidateSymbol(symbol)
	if err != nil {
		return
	}
	err = ValidateInterval(interval)
	if err != nil {
		return
	}
	err = ValidateDates(startDate, endDate)
	if err != nil {
		return
	}
	err = ValidateTimeZone(tz)
	if err != nil {
		return
	}
	err = ValidateOutputSize(outputSize)

	return
}

type Bar struct {
	Datetime string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
}

type BarString struct {
	Datetime string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
}

func ParseValuesToFloat(inputValues []BarString) (values []Bar) {
	for _, v := range inputValues {
		value := Bar{
			Datetime: v.Datetime,
		}

		open, _ := strconv.ParseFloat(v.Open, 64)
		close, _ := strconv.ParseFloat(v.Close, 64)
		high, _ := strconv.ParseFloat(v.High, 64)
		low, _ := strconv.ParseFloat(v.Low, 64)
		volume, _ := strconv.Atoi(v.Volume)

		value.Open = open
		value.Close = close
		value.High = high
		value.Low = low
		value.Volume = volume

		values = append(values, value)
	}
	return
}
