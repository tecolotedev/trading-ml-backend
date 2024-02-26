package indicators

import (
	"github.com/tecolotedev/trading-ml-backend/utils"
)

type SMAInput struct {
	SeriesType string // open, hight, low, close
	Values     []utils.Bar
	FillNA     string
	Periods    int
}

type SMAOuput struct {
	Datetime string
	Value    float64
}

type MakeSMAInput struct {
	Datetime string
	value    float64
}

func MakeSMA(input []MakeSMAInput, fillNA string, periods int) (output []SMAOuput) {
	amount := 0.0

	for i := 0; i < len(input); i++ {
		if i+1 <= periods {
			o := SMAOuput{
				Datetime: input[i].Datetime,
				Value:    0,
			}
			output = append(output, o)
			amount += input[i].value
		} else {
			o := SMAOuput{
				Datetime: input[i].Datetime,
				Value:    amount / float64(periods),
			}
			output = append(output, o)
			amount += input[i].value
			amount -= input[i-periods].value
		}
	}
	return
}

func GetSMA(input SMAInput) (output []SMAOuput) {
	inputValues := []MakeSMAInput{}

	switch input.SeriesType {

	case "open":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeSMAInput{Datetime: v.Datetime, value: v.Open})
		}
		return MakeSMA(inputValues, input.FillNA, input.Periods)
	case "high":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeSMAInput{Datetime: v.Datetime, value: v.High})
		}
		return MakeSMA(inputValues, input.FillNA, input.Periods)
	case "low":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeSMAInput{Datetime: v.Datetime, value: v.Low})
		}
		return MakeSMA(inputValues, input.FillNA, input.Periods)
	case "close":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeSMAInput{Datetime: v.Datetime, value: v.Close})
		}
		return MakeSMA(inputValues, input.FillNA, input.Periods)
	}
	return
}
