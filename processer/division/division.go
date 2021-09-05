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

func (p division) Calculate(first float64, last float64) (float64, error) {
	if last == 0 {
		return 0, errors.New("division error, last number can't be 0")
	}
	return math.Round(first/last*math.Pow(10, p.precision)) / math.Pow(10, p.precision), nil
}
