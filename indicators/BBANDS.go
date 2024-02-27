package indicators

import (
	"math"

	"github.com/tecolotedev/trading-ml-backend/utils"
)

type BBANDSInput struct {
	SeriesType string // open, hight, low, close
	Values     []utils.Bar
	FillNA     string
	Periods    int
	SD         int
}

type BBANDSOuput struct {
	Datetime   string
	SMA        float64
	UPPER_BAND float64
	LOWER_BAND float64
}

func MakeBBANDS(input []MakeIndicatorInput, fillNA string, periods, sd int) (output []BBANDSOuput) {
	sma := MakeSMA(input, fillNA, periods)

	sum := 0.0

	for i := 0; i < len(input); i++ {
		if i+1 <= periods {
			sum += input[i].value
			o := BBANDSOuput{
				Datetime:   input[i].Datetime,
				SMA:        sma[i].Value,
				UPPER_BAND: 0,
				LOWER_BAND: 0,
			}
			output = append(output, o)
		} else {
			mean := sum / float64(periods)
			prices := input[i-periods : i]

			sqrSum := 0.0

			for j := 0; j < len(prices); j++ {
				sqrSum += math.Pow((prices[j].value - mean), 2)
			}

			sd := math.Sqrt(sqrSum / float64(periods))

			o := BBANDSOuput{
				Datetime:   input[i].Datetime,
				SMA:        sma[i].Value,
				UPPER_BAND: sma[i].Value + sd,
				LOWER_BAND: sma[i].Value - sd,
			}
			output = append(output, o)

			sum += sma[i].Value
			sum -= sma[i-periods].Value

		}
	}

	return
}

func GetBBANDS(input BBANDSInput) (output []BBANDSOuput) {
	inputValues := []MakeIndicatorInput{}

	switch input.SeriesType {

	case "open":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Open})
		}
		return MakeBBANDS(inputValues, input.FillNA, input.Periods, input.SD)
	case "high":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.High})
		}
		return MakeBBANDS(inputValues, input.FillNA, input.Periods, input.SD)
	case "low":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Low})
		}
		return MakeBBANDS(inputValues, input.FillNA, input.Periods, input.SD)
	case "close":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Close})
		}
		return MakeBBANDS(inputValues, input.FillNA, input.Periods, input.SD)
	}
	return
}
