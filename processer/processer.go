package processer

import (
	"errors"

	"github.com/ffo32167/calculator/processer/division"
	"github.com/ffo32167/calculator/processer/substraction"
)

type AdditionalData struct {
	Precision int
}

type Action interface {
	Calculate(float64, float64) (float64, error)
}

func PickAction(action string) (Action, error) {
	switch action {
	case "division":
		return division.NewDivision(), nil
	case "substraction":
		return substraction.NewSubstraction(), nil
	default:
		return nil, errors.New("cant PickAction, unknown action")
	}
}
