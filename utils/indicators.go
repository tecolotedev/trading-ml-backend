package utils

import (
	"fmt"
	"strconv"
)

func ValidateTimePeriod(tp int) (err error) {
	if tp > 1 && tp < 800 {
		return
	}

	return fmt.Errorf("time_period must be between 1 and 800 including")
}

var maTypes = map[string]string{
	"SMA":   "SMA",
	"EMA":   "EMA",
	"WMA":   "WMA",
	"DEMA":  "DEMA",
	"TEMA":  "TEMA",
	"TRIMA": "TRIMA",
	"KAMA":  "KAMA",
	"MAMA":  "MAMA",
	"T3MA":  "T3MA",
}

func ValidateMAType(maType string) (err error) {
	for key := range maTypes {
		if key == maType {
			return
		}
	}

	return fmt.Errorf("ma_type is invalid")
}

var seriesTypes = map[string]string{
	"open":  "open",
	"close": "close",
	"high":  "high",
	"low":   "low",
}

func ValidateSeriesType(seriesType string) (err error) {
	for key := range seriesTypes {
		if key == seriesType {
			return
		}
	}

	return fmt.Errorf("series_type is invalid")
}

func ValidateMAParams(timePeriod int, maType, seriesType string) (err error) {
	err = ValidateTimePeriod(timePeriod)
	if err != nil {
		return
	}
	err = ValidateMAType(maType)
	if err != nil {
		return
	}
	err = ValidateSeriesType(seriesType)
	return
}

type MAValueResponse struct {
	Datetime string
	MA       string
}
type MAValueParsed struct {
	Datetime string
	MA       float64
}

func ParseMAValues(inputValues []MAValueResponse) (values []MAValueParsed) {
	for _, v := range inputValues {
		value := MAValueParsed{
			Datetime: v.Datetime,
		}

		ma, _ := strconv.ParseFloat(v.MA, 64)
		value.MA = ma

		values = append(values, value)
	}
	return
}
