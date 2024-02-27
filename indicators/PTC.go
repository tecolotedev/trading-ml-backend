package indicators

import (
	"math"

	"github.com/tecolotedev/trading-ml-backend/utils"
)

type PTCInput struct {
	SeriesType string // open, hight, low, close
	Values     []utils.Bar
	FillNA     string
}

type PTCOuput struct {
	Datetime string
	Value    float64
}

func MakePTC(input []MakeIndicatorInput, fillNA string) (output []PTCOuput) {

	for i := 0; i < len(input); i++ {
		if i == 0 {
			o := PTCOuput{
				Datetime: input[i].Datetime,
			}
			output = append(output, o)
		} else {
			o := PTCOuput{
				Datetime: input[i].Datetime,
				Value:    math.Log(input[i].value / input[i-1].value),
			}
			output = append(output, o)
		}
	}

	return
}

func GetPTC(input PTCInput) (output []PTCOuput) {
	inputValues := []MakeIndicatorInput{}

	switch input.SeriesType {

	case "open":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Open})
		}
		return MakePTC(inputValues, input.FillNA)
	case "high":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.High})
		}
		return MakePTC(inputValues, input.FillNA)
	case "low":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Low})
		}
		return MakePTC(inputValues, input.FillNA)
	case "close":
		for _, v := range input.Values {
			inputValues = append(inputValues, MakeIndicatorInput{Datetime: v.Datetime, value: v.Close})
		}
		return MakePTC(inputValues, input.FillNA)
	}
	return
}
