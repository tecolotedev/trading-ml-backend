package utils

import (
	"fmt"
	"strconv"
)

/*
 * Utils for Moving Average indicator
 */
func ValidatePeriod(tp int, namePeriod string) (err error) {
	if tp >= 1 && tp <= 800 {
		return
	}

	return fmt.Errorf("%s must be between 1 and 800 including", namePeriod)
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
	err = ValidatePeriod(timePeriod, "time_period")
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
	Ma       string
}
type MAValueParsed struct {
	Datetime string
	Ma       float64
}

func ParseMAValues(inputValues []MAValueResponse) (values []MAValueParsed) {
	for _, v := range inputValues {
		value := MAValueParsed{
			Datetime: v.Datetime,
		}

		ma, _ := strconv.ParseFloat(v.Ma, 64)
		value.Ma = ma

		values = append(values, value)
	}
	return
}

/*
 * Utils for Moving Average Convergence Divergence (MACD) indicator
 */

type MACDValueResponse struct {
	Datetime    string
	MACD        string
	MACD_SIGNAL string
	MACD_HIST   string
}
type MACDValueParsed struct {
	Datetime    string
	MACD        float64
	MACD_SIGNAL float64
	MACD_HIST   float64
}

func ParseMACDValues(inputValues []MACDValueResponse) (values []MACDValueParsed) {
	for _, v := range inputValues {
		value := MACDValueParsed{
			Datetime: v.Datetime,
		}

		macd, _ := strconv.ParseFloat(v.MACD, 64)
		macd_signal, _ := strconv.ParseFloat(v.MACD_SIGNAL, 64)
		macd_hist, _ := strconv.ParseFloat(v.MACD_HIST, 64)

		value.MACD = macd
		value.MACD_SIGNAL = macd_signal
		value.MACD_HIST = macd_hist

		values = append(values, value)
	}
	return
}

func ValidateMACDParams(fastPeriod, signalPeriod, slowPeriod int, seriesType string) (err error) {
	err = ValidatePeriod(fastPeriod, "fast_period")
	if err != nil {
		return
	}
	err = ValidatePeriod(signalPeriod, "signal_period")
	if err != nil {
		return
	}
	err = ValidatePeriod(slowPeriod, "slow_period")
	if err != nil {
		return
	}
	err = ValidateSeriesType(seriesType)
	return
}

/*
 * Utils for Relative Strength Index (RSI) indicator
 */

type RSIValueResponse struct {
	Datetime string
	RSI      string
}
type RSIValueParsed struct {
	Datetime string
	RSI      float64
}

func ParseRSIValues(inputValues []RSIValueResponse) (values []RSIValueParsed) {
	for _, v := range inputValues {
		value := RSIValueParsed{
			Datetime: v.Datetime,
		}

		rsi, _ := strconv.ParseFloat(v.RSI, 64)
		value.RSI = rsi

		values = append(values, value)
	}
	return
}

func ValidateRSIParams(timePeriod int) (err error) {
	err = ValidatePeriod(timePeriod, "time_period")
	return
}

/*
 * Utils for Bollinger Bands (BBANDS) indicator
 */

type BBANDSValueResponse struct {
	Datetime    string
	UPPER_BAND  string
	MIDDLE_BAND string
	LOWER_BAND  string
}
type BBANDSValueParsed struct {
	Datetime    string
	UPPER_BAND  float64
	MIDDLE_BAND float64
	LOWER_BAND  float64
}

func ParseBBANDSValues(inputValues []BBANDSValueResponse) (values []BBANDSValueParsed) {
	for _, v := range inputValues {
		value := BBANDSValueParsed{
			Datetime: v.Datetime,
		}

		uBand, _ := strconv.ParseFloat(v.UPPER_BAND, 64)
		mBand, _ := strconv.ParseFloat(v.MIDDLE_BAND, 64)
		lBand, _ := strconv.ParseFloat(v.LOWER_BAND, 64)

		value.UPPER_BAND = uBand
		value.MIDDLE_BAND = mBand
		value.LOWER_BAND = lBand

		values = append(values, value)
	}
	return
}

func ValidateBBANDSParams(timePeriod int, maType, seriesType string) (err error) {
	err = ValidatePeriod(timePeriod, "time_period")
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
