package division

import (
	"errors"
	"math"
)

type division struct {
	precision float64
}

func NewDivision(precision float64) division {
	return division{precision: precision}
}

func (d division) Calculate(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Calculate division error, last number can't be 0")
	}
	return round(num1/num2, d.precision), nil
}

func round(in float64, precision float64) float64 {
	return math.Round(in*math.Pow(10, precision)) / math.Pow(10, precision)
}
