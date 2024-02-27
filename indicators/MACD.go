package indicators

import (
	"github.com/tecolotedev/trading-ml-backend/utils"
)

type MACDInput struct {
	SeriesType  string // open, hight, low, close
	Values      []utils.Bar
	FillNA      string
	FastPeriods int
	SlowPeriods int
	MACDPeriods int
}

type MACDOuput struct {
	Datetime string
	SlowEMA  float64
	FastEMA  float64
	MACD     float64
	MACD_SMA float64
}

func MakeMACD(input []MakeIndicatorInput, fillNA string, fastPeriods, slowPeriods, macdPeriods int) (output []MACDOuput) {
	slowEMA := MakeEMA(input, fillNA, slowPeriods)
	fastEMA := MakeEMA(input, fillNA, fastPeriods)

	_macdInput := []MakeIndicatorInput{}

	for i := 0; i < len(input); i++ {
		if slowEMA[i].Value == 0 || fastEMA[i].Value == 0 {
			_macdInput = append(_macdInput, MakeIndicatorInput{value: 0.0})
		} else {
			v := slowEMA[i].Value - fastEMA[i].Value
			_macdInput = append(_macdInput, MakeIndicatorInput{value: v})
		}
	}

	macdSMA := MakeEMA(_macdInput, fillNA, macdPeriods)

	for i := 0; i < len(input); i++ {
		o := MACDOuput{
			Datetime: input[i].Datetime,
			SlowEMA:  slowEMA[i].Value,
			FastEMA:  fastEMA[i].Value,
			MACD:     _macdInput[i].value,
			MACD_SMA: macdSMA[i].Value,
		}
		output = append(output, o)
	}

	return
}

func GetMACD(input MACDInput) (output []MACDOuput) {
	inputValues := []MakeIndicatorInput{}

	switch input.SeriesType {

	case "open":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Open})
		}
		return MakeMACD(inputValues, input.FillNA, input.FastPeriods, input.SlowPeriods, input.MACDPeriods)
	case "high":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.High})
		}
		return MakeMACD(inputValues, input.FillNA, input.FastPeriods, input.SlowPeriods, input.MACDPeriods)
	case "low":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Low})
		}
		return MakeMACD(inputValues, input.FillNA, input.FastPeriods, input.SlowPeriods, input.MACDPeriods)
	case "close":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Close})
		}
		return MakeMACD(inputValues, input.FillNA, input.FastPeriods, input.SlowPeriods, input.MACDPeriods)
	}
	return
}
