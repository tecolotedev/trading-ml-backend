package indicators

import (
	"math"

	"github.com/tecolotedev/trading-ml-backend/utils"
)

type RSIInput struct {
	SeriesType string // open, hight, low, close
	Values     []utils.Bar
	FillNA     string
	Periods    int
}

type RSIOuput struct {
	Datetime string
	Value    float64
}

func MakeRSI(input []MakeIndicatorInput, fillNA string, periods int) (output []RSIOuput) {

	totalLoss := 0.0
	totalGain := 0.0

	avgLoss := 0.0
	avgGain := 0.0

	for i := 0; i < len(input); i++ {
		if i == 0 {
			o := RSIOuput{
				Datetime: input[i].Datetime,
				Value:    0,
			}
			output = append(output, o)
			// skip first row
			continue
		}

		change := input[i].value - input[i-1].value

		if i+1 <= periods {
			if change > 0 {
				totalGain += change
			} else {
				totalLoss += math.Abs(change)
			}
			o := RSIOuput{
				Datetime: input[i].Datetime,
				Value:    0,
			}
			output = append(output, o)
		} else if i == periods {
			avgLoss = totalLoss / float64(periods)
			avgGain = totalGain / float64(periods)

			rs := avgGain / avgLoss
			rsi := 100 - (100 / (1 + rs))

			o := RSIOuput{
				Datetime: input[i].Datetime,
				Value:    rsi,
			}
			output = append(output, o)

		} else {
			if change > 0 {
				avgLoss = (avgLoss * (float64(periods) - 1)) / float64(periods)
				avgGain = (avgGain*(float64(periods)-1) + change) / float64(periods)
			} else {
				avgLoss = (avgLoss*(float64(periods)-1) + math.Abs(change)) / float64(periods)
				avgGain = (avgGain * (float64(periods) - 1)) / float64(periods)
			}

			rs := avgGain / avgLoss
			rsi := 100 - (100 / (1 + rs))

			o := RSIOuput{
				Datetime: input[i].Datetime,
				Value:    rsi,
			}
			output = append(output, o)
		}

	}
	return
}

func GetRSI(input RSIInput) (output []RSIOuput) {
	inputValues := []MakeIndicatorInput{}

	switch input.SeriesType {

	case "open":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Open})
		}
		return MakeRSI(inputValues, input.FillNA, input.Periods)
	case "high":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.High})
		}
		return MakeRSI(inputValues, input.FillNA, input.Periods)
	case "low":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Low})
		}
		return MakeRSI(inputValues, input.FillNA, input.Periods)
	case "close":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Close})
		}
		return MakeRSI(inputValues, input.FillNA, input.Periods)
	}
	return
}
