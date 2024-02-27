package indicators

import (
	"github.com/tecolotedev/trading-ml-backend/utils"
)

type EMAInput struct {
	SeriesType string // open, hight, low, close
	Values     []utils.Bar
	FillNA     string
	Periods    int
}

type EMAOuput struct {
	Datetime string
	Value    float64
}

func MakeEMA(input []MakeIndicatorInput, fillNA string, periods int) (output []EMAOuput) {
	lastEma := 0.0
	amount := 0.0
	alpha := 2.0 / float64((periods + 1))

	for i := 0; i < len(input); i++ {
		if i+1 <= periods {
			o := EMAOuput{
				Datetime: input[i].Datetime,
				Value:    0,
			}
			output = append(output, o)
			amount += input[i].value
		} else if i+1 == periods+1 {
			// first value of periods is like SMA

			ema := alpha*input[i].value + (1-alpha)*(amount/float64(periods))
			o := EMAOuput{
				Datetime: input[i].Datetime,
				Value:    ema,
			}
			output = append(output, o)
			lastEma = ema

		} else {
			ema := alpha*input[i].value + (1-alpha)*lastEma
			o := EMAOuput{
				Datetime: input[i].Datetime,
				Value:    ema,
			}
			output = append(output, o)
			lastEma = ema

		}
	}
	return
}

func GetEMA(input EMAInput) (output []EMAOuput) {
	inputValues := []MakeIndicatorInput{}

	switch input.SeriesType {

	case "open":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Open})
		}
		return MakeEMA(inputValues, input.FillNA, input.Periods)
	case "high":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.High})
		}
		return MakeEMA(inputValues, input.FillNA, input.Periods)
	case "low":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Low})
		}
		return MakeEMA(inputValues, input.FillNA, input.Periods)
	case "close":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Close})
		}
		return MakeEMA(inputValues, input.FillNA, input.Periods)
	}
	return
}
