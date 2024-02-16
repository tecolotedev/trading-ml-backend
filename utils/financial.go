package utils

import (
	"strconv"
)

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

func ValidateSymbol(symbol string) bool {
	isValid := false

	for key := range Symbols {
		if key == symbol {
			isValid = true
		}
	}

	return isValid
}

func ValidateInterval(interval string) bool {
	isValid := false

	for key := range Intervals {
		if key == interval {
			isValid = true
		}
	}

	return isValid
}

type ValueParsed struct {
	Datetime string  `json:"datetime"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   int     `json:"volume"`
}

type ValueResponse struct {
	Datetime string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
}

func ParseValuesToFloat(inputValues []ValueResponse) (values []ValueParsed) {
	for _, v := range inputValues {
		value := ValueParsed{
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

func ValidateOutputSize(outputsize int) (isValid bool) {
	if outputsize > 0 && outputsize <= 5000 {
		isValid = true
	}
	return
}
