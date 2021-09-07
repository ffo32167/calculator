package division

import (
	"errors"
	"math"
)

type division struct {
	precision float64
}

func NewDivision() division {
	return division{precision: 2}
}

func round(in float64, precision float64) float64 {
	return math.Round(in*math.Pow(10, precision)) / math.Pow(10, precision)
}

func (d division) Calculate(first float64, last float64) (float64, error) {
	if last == 0 {
		return 0, errors.New("Calculate division error, last number can't be 0")
	}
	return round(first/last, d.precision), nil
}
