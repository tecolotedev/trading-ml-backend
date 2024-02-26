package utils

func SumInt(numbers []int) (amount int) {
	for i := 0; i < len(numbers); i++ {
		amount += numbers[i]
	}

	return
}

func SumFloat64(numbers []float64) (amount float64) {

	for i := 0; i < len(numbers); i++ {
		amount += numbers[i]
	}

	return

}
