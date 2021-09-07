package processer

import (
	"errors"

	"github.com/ffo32167/calculator/processer/division"
	"github.com/ffo32167/calculator/processer/substraction"
)

type actionData struct {
	Action
	AdditionalData
}

type AdditionalData struct {
	Precision float64
}

type Action interface {
	Calculate(num1, num2 float64) (float64, error)
}

func NewActionData(a Action, ad AdditionalData) actionData {
	return actionData{a, ad}
}

func PickAction(action string, ad AdditionalData) (Action, error) {
	switch action {
	case "division":
		return division.NewDivision(ad.Precision), nil
	case "substraction":
		return substraction.NewSubstraction(), nil
	default:
		return nil, errors.New("cant PickAction, unknown action")
	}
}
